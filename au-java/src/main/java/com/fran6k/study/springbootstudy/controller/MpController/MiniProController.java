package com.fran6k.study.springbootstudy.controller.MpController;


import com.alibaba.fastjson.JSONObject;
import com.fran6k.study.springbootstudy.bean.WxBean.Wxuser;
import com.fran6k.study.springbootstudy.dao.AssRepository;
import com.fran6k.study.springbootstudy.dao.WxuserRepository;
import com.fran6k.study.springbootstudy.utils.JwtUtil;
import com.fran6k.study.springbootstudy.utils.ResultModel;
import com.fran6k.study.springbootstudy.utils.UIMLoginUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.MediaType;
import org.springframework.http.converter.StringHttpMessageConverter;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

import java.io.IOException;
import java.nio.charset.Charset;
import java.util.Map;

//小程序登录接口
public class MiniProController {
    @Autowired
    private WxuserRepository wxuserRepository;
    @Autowired
    private AssRepository assRepository;

    @Value("${mpconfig.appid}")
    private String appid;
    @Value("${mpconfig.appsecret}")
    private String appsecret;


    @GetMapping("mpapi/mytest")
    public String test() throws IOException{
        String name = "cl";
        return name;
    }

    /*
     * 获取用户信息
     *
     * */
    @GetMapping("mpapi/getUserInfo")
    @ResponseBody
    public ResultModel getUserInfo(@RequestParam("code") String code) throws IOException{
        //第一步，通过code换取appid和appsecret
        String urlOne = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + appid +
                "&secret=" + appsecret +
                "&code=" + code +
                "&grant_type=authorization_code";

        RestTemplate restTemplate=new RestTemplate();
        // 解决乱码问题
        restTemplate.getMessageConverters().set(1,new StringHttpMessageConverter(Charset.forName("UTF-8")));
        // 请求体
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_FORM_URLENCODED);
        HttpEntity<String> formEntity = new HttpEntity<String>(null,headers);
        // 第二步，得到openid和access_token
        String body= restTemplate.exchange(urlOne , HttpMethod.GET, formEntity, String.class).getBody();
        JSONObject jsonObject = JSONObject.parseObject(body);
        System.out.println("测试"+body);
        String openid = jsonObject.getString("openid");
        String access_token = jsonObject.getString("access_token");
        String errcode = jsonObject.getString("errcode");
        if (errcode!=null){
            return new ResultModel(20003,"无效code");
        }
        // 查询数据库判断是否有此人
        Wxuser wxuser = wxuserRepository.findByOpenId(openid);
        // 有，返回信息和token
        if(wxuser!=null){
            JSONObject wxuserInfo = new JSONObject();
            wxuserInfo.put("avatar",wxuser.getAvatar());
            wxuserInfo.put("nickName",wxuser.getNickname());
            wxuserInfo.put("studentId",wxuser.getStudentId());
            wxuserInfo.put("realName",wxuser.getRealName());
            wxuserInfo.put("assName", wxuser.getAssObject());
            // 签发token
            String token = JwtUtil.sign(openid,"-1");
            wxuserInfo.put("token",token);
            ResultModel resultModel = new ResultModel(20000,"ok",wxuserInfo);
            System.out.println(resultModel);
            return resultModel;
        }
        // 没有，返回代码20001跳转页面，+token
        else{
            // 第三步：拉取用户信息(需scope为 snsapi_userinfo)
            String urlTwo = "https://api.weixin.qq.com/sns/userinfo?access_token=" + access_token +
                    "&openid=" + openid +
                    "&lang=zh_CN";

            String userInfo= restTemplate.exchange(urlTwo , HttpMethod.GET, formEntity, String.class).getBody();
            JSONObject userInfoJson = JSONObject.parseObject(userInfo);
            userInfoJson.put("openid",openid);
            ResultModel resultModel = new ResultModel(20001,"ok",userInfoJson);
            return resultModel;
        }
    }

    /*
     * 绑定学号接口
     * 传入账号和密码
     * */
    @PostMapping("mpapi/addWxUser")
    @ResponseBody
    public ResultModel addWxUser(@RequestBody Map params) throws IOException {
        String studentId = params.get("studentId").toString();
        String password = params.get("password").toString();
        JSONObject UIMJSON = UIMLoginUtil.veri(studentId, password);
        //验证成功
        if (UIMJSON.getString("code").equals("20000"));
        else {
            ResultModel resultModel = new ResultModel(20002, "学号密码验证失败");
            return resultModel;
        }
        //验证之后查看是否已经验证
        Wxuser wxuser = wxuserRepository.findByStudentId(studentId);
        if (wxuser != null) {
            ResultModel resultModel = new ResultModel(20000, "您已经绑定过学号，无需再次绑定");
            return resultModel;
        }
        //没有验证那么将用户信息存入数据库
        wxuser = new Wxuser();
        wxuser.setOpenId(params.get("openId").toString());
        wxuser.setStudentId(params.get("studentId").toString());
        wxuser.setRealName(params.get("realName").toString());
        wxuser.setWeChatId(params.get("weChatId").toString());
        wxuser.setPhoneNum(params.get("phoneNum").toString());
        wxuser.setNickname(params.get("nickName").toString());
        wxuser.setAvatar(params.get("avatar").toString());
//        System.out.println(params.toString());
        wxuserRepository.save(wxuser);
        ResultModel resultModel = new ResultModel(20000, "绑定成功！");
        return resultModel;
    }

}
