package com.fran6k.study.springbootstudy.controller.WxController;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.fran6k.study.springbootstudy.bean.Ass;
import com.fran6k.study.springbootstudy.bean.User;
import com.fran6k.study.springbootstudy.bean.WxBean.Wxuser;
import com.fran6k.study.springbootstudy.controller.AdminController;
import com.fran6k.study.springbootstudy.dao.AssRepository;
import com.fran6k.study.springbootstudy.dao.WxuserRepository;
import com.fran6k.study.springbootstudy.utils.*;
import net.sf.json.JSONArray;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.data.domain.Sort;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpMethod;
import org.springframework.http.MediaType;
import org.springframework.http.converter.StringHttpMessageConverter;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

import java.io.IOException;
import java.nio.charset.Charset;
import java.util.List;
import java.util.Map;

@Controller
public class WxApiController {
    @Autowired
    private WxuserRepository wxuserRepository;
    @Autowired
    private AssRepository assRepository;

    @Value("${wxconfig.appid}")
    private String appid;
    @Value("${wxconfig.appsecret}")
    private String appsecret;

    //社团列表
    private JSONArray assList(List<Ass> associations){
        JSONArray jsonArray = new JSONArray();
        for(Ass s : associations){
            JSONObject jsonObject = new JSONObject();
            jsonObject.put("assId",s.getAssid());
            jsonObject.put("assName",s.getAssname());
            jsonObject.put("assType",s.getAsstype());
            jsonObject.put("logo",s.getLogo());
            jsonObject.put("description",s.getAssDescription());
            jsonArray.add(jsonObject);
        }
        return jsonArray;
    }

    @GetMapping("wxapi/getUserInfo")
    @ResponseBody
    public ResultModel getUserInfo(@RequestParam("code") String code) throws IOException {

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
        // 有，返回信息和toke
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
    @PostMapping("wxapi/addWxUser")
    @ResponseBody
    public ResultModel addWxUser(@RequestBody Map params) throws IOException {
        String studentId = params.get("studentId").toString();
        String password = params.get("password").toString();
        JSONObject UIMJSON = UIMLoginUtil.veri(studentId,password);

        if (UIMJSON.getString("code").equals("20000"));
        else {
            ResultModel resultModel = new ResultModel(20002, "学号密码验证失败");
            return resultModel;
        }
        Wxuser wxuser = wxuserRepository.findByStudentId(studentId);
        if (wxuser != null) {
            ResultModel resultModel = new ResultModel(20000, "您已经绑定过学号，无需再次绑定");
            return resultModel;
        }
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

    @GetMapping("wxapi/getTotalAssociation")
    @ResponseBody
    public ResultModelArrary getTotalAssociation() throws IOException {
        JSONArray jsonArray = new JSONArray();
        List<Ass> associations= assRepository.findAll();
        ResultModelArrary resultModel = new ResultModelArrary(20000, "",assList(associations));
        return resultModel;
    }
    @GetMapping("wxapi/getExecllentAssociation")
    @ResponseBody
    public ResultModelArrary getExecllentAssociation() throws IOException {
        List<Ass> associations= assRepository.findAllByIsExecllent(1);
        ResultModelArrary resultModel = new ResultModelArrary(20000, "",assList(associations));
        return resultModel;
    }
    @GetMapping("wxapi/test")
    @ResponseBody
    public String test() throws IOException {
        JSONObject jsonObject = UIMLoginUtil.veri("201904020217","sztu@180016");
        return null;
    }
    @GetMapping("wxapi/getAssociationsNameMapAssid")
    @ResponseBody
    public ResultModel getAssociationsNameMapAssid() throws IOException {
        JSONObject jsonObject = new JSONObject();
        List<Ass> associations= assRepository.findAllByAssidIsNot(1);
        for(Ass s : associations){
            jsonObject.put(String.valueOf(s.getAssid()),s.getAssname());
        }
        ResultModel resultModel = new ResultModel(20000, "",jsonObject);
        return resultModel;
    }
    @GetMapping("wxapi/getAssociationsByAssid")
    @ResponseBody
    public ResultModel getAssociationsByAssid(@RequestParam("assId")int assId) throws IOException {
        Ass ass = assRepository.findByAssid(assId);
        JSONObject jsonObject = ass.getAss();
        return new ResultModel(20000, "",jsonObject);
    }
    @GetMapping("wxapi/getTwitterPic")
    @ResponseBody
    public ResultModel getTwitterPic(@RequestParam("url")String url) throws IOException {
        RestTemplate restTemplate = new RestTemplate();
        String body = restTemplate.getForEntity(url,String.class).getBody();
        assert body != null;
        //解析封面图片地址
        int begin = body.indexOf("var msg_cdn_url");
        int end = body.indexOf("var cdn_url_1_1");
        String picUrl = body.substring(begin,end);
        picUrl = picUrl.split("\"")[1];
        //解析推文名称
        int title_first = body.lastIndexOf("<meta property=\"twitter:title\" content=\"");
        int title_last = body.indexOf("<meta property=\"twitter:creator\"");
        String title = body.substring(title_first,title_last).split("\"")[3];
        //解析时间
        int time_first = body.indexOf("if(window.__second_open__)return;");
        int time_last = body.indexOf("e(t,n,i,document.getElementById(\"publish_time\"));");
        String time = body.substring(time_first,time_last).split("\"")[5];
        JSONObject jsonObject = new JSONObject();
        jsonObject.put("picUrl",picUrl);
        jsonObject.put("title",title);
        jsonObject.put("time",time);
        return new ResultModel(20000, "",jsonObject);
    }
}
