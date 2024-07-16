package com.fran6k.study.springbootstudy.controller.WxController;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.fran6k.study.springbootstudy.bean.Ass;
import com.fran6k.study.springbootstudy.bean.WxBean.InterViewAss;
import com.fran6k.study.springbootstudy.bean.WxBean.InterViewEnum;
import com.fran6k.study.springbootstudy.bean.WxBean.InterViewUser;
import com.fran6k.study.springbootstudy.bean.WxBean.Wxuser;
import com.fran6k.study.springbootstudy.dao.AssRepository;
import com.fran6k.study.springbootstudy.dao.InterViewAssRepository;
import com.fran6k.study.springbootstudy.dao.InterViewUserRepository;
import com.fran6k.study.springbootstudy.dao.WxuserRepository;
import com.fran6k.study.springbootstudy.utils.CheckTime;
import com.fran6k.study.springbootstudy.utils.ResultModel;
import com.fran6k.study.springbootstudy.utils.ResultModelArrary;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.core.io.FileSystemResource;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.converter.StringHttpMessageConverter;
import org.springframework.stereotype.Controller;
import org.springframework.util.LinkedMultiValueMap;
import org.springframework.util.MultiValueMap;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.multipart.MultipartFile;

import javax.servlet.http.HttpServletRequest;
import java.io.*;
import java.net.URLDecoder;
import java.net.URLEncoder;
import java.nio.charset.Charset;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.*;

@Controller
public class WxInterViewController {
    @Autowired
    InterViewAssRepository interViewAssRepository;
    @Autowired
    InterViewUserRepository interViewUserRepository;
    @Autowired
    WxuserRepository wxuserRepository;
    @Autowired
    AssRepository assRepository;

    @Value("${interviewass.opentime}")
    private String openTime;
    @Value("${interviewass.endtime}")
    private String endTime;

    /**
     * 添加面试用户
     *
     * @return
     * @throws IOException
     */
    @PostMapping("wxapi/interview/addInterviewUser")
    @ResponseBody
    public ResultModel addInterviewUser(@RequestBody Map params) throws IOException {
        ResultModel resultModel = new ResultModel(20000, "");
        JSONObject jsonObject = new JSONObject();

        String studentId = params.get("studentId").toString();
        String description = params.get("description").toString();
        String name = params.get("name").toString();
        String sex = params.get("sex").toString();
        String phoneNum = params.get("phoneNum").toString();
        String wxNum = params.get("wxNum").toString();
        int assId = Integer.parseInt(params.get("assId").toString());
        //1.没有姓名或者没有学号的错误
        if (name == null || studentId == null) {
            jsonObject.put("status", 0);
            resultModel.setMsg("出现未知错误,请重新进入系统");
            resultModel.setData(jsonObject);
            return resultModel;
        }
        //2.检测时间

        if (!CheckTime.checkTime(openTime, endTime)) {
            jsonObject.put("status", 0);
            resultModel.setMsg("报名时间为:" + openTime + "到" + endTime);
            resultModel.setData(jsonObject);
            return resultModel;
        }
        //2.社团已满的错误
        Wxuser wxuser = wxuserRepository.findByStudentId(studentId);
        Ass ass1 = null;
        Ass ass2 = null;
        try {
            ass1 = wxuser.getFirstAss();
            ass2 = wxuser.getSecondAss();
        } catch (Exception e) {
        }
        if (ass1 != null && ass2 != null) {
            jsonObject.put("status", 0);
            resultModel.setMsg("您已经加入了2个社团哦");
            resultModel.setData(jsonObject);
            return resultModel;
        }
        //3.重复提交报名的错误
        InterViewAss interViewAss = interViewAssRepository.findByAssId(assId);
        InterViewUser interViewUser = interViewUserRepository.findByStudentIdAndInterViewAss(studentId, interViewAss);
        if (interViewUser != null) {
            jsonObject.put("status", 0);
            resultModel.setData(jsonObject);
            resultModel.setMsg("您已经报名了此社团，请勿重复提交");
            return resultModel;
        }
        if ((ass1 != null && ass1.getAssid() == interViewAss.getAssId()) || (ass2 != null && ass2.getAssid() == interViewAss.getAssId())) {
            jsonObject.put("status", 0);
            resultModel.setData(jsonObject);
            resultModel.setMsg("您已经加入此社团了");
            return resultModel;
        }

            //4.报名数限制 2
        ArrayList<InterViewUser> interViewUsers = interViewUserRepository.findAllByStudentId(studentId);
        if (interViewUsers.size() == 2) {
            jsonObject.put("status", 0);
            resultModel.setData(jsonObject);
            resultModel.setMsg("最多只可以报名2个社团哦");
            return resultModel;
        }
        interViewUser = new InterViewUser(studentId, name, sex, description, interViewAss, phoneNum, wxNum);
        //interViewUser.setButtonControl(1);默认为1，使用权在社团
        interViewUser.setBackMessage(interViewAss.getShowMessage());
        interViewUserRepository.save(interViewUser);
        jsonObject.put("status", 1);
        resultModel.setMsg("报名成功!");
        resultModel.setData(jsonObject);
        return resultModel;
    }

    /**
     * 学生确认加入社团
     *
     * @param id
     * @return
     */
    @GetMapping("wxapi/interview/studentConfirm")
    @ResponseBody
    public ResultModel studentConfirm(@RequestParam("id") String id) {
        JSONObject jsonObject = new JSONObject();
        ResultModel resultModel = new ResultModel(20000, "");
        InterViewUser interViewUser = interViewUserRepository.getOne(id);
        InterViewEnum interViewEnum = interViewUser.getInterViewStatus();
        Wxuser wxuser = wxuserRepository.findByStudentId(interViewUser.getStudentId());
        //检测社团是否已满
        if (wxuser.getFirstAss() != null && wxuser.getSecondAss() != null) {
            jsonObject.put("status", 0);
            resultModel.setData(jsonObject);
            resultModel.setMsg("最只能加入2个社团");
            return resultModel;
        }
        //判断状态是否正确
        if (interViewEnum == InterViewEnum.Stage_Two_Success) {
            //更改并更新面试表中的状态
            interViewEnum = InterViewEnum.Stage_Three_Success;
            interViewUser.setInterViewStatus(interViewEnum);
            //设置控制按钮
            interViewUser.setButtonControl(0);
            //设置成功消息+拼接图片二维码
            interViewUser.setBackMessage(interViewUser.getInterViewAss().getCodeUrl() + "|" + interViewUser.getInterViewAss().getConfirmJoinMessage());
            interViewUserRepository.save(interViewUser);
            //设置学生所在社团
            if (wxuser.getFirstAss() == null) {
                wxuser.setFirstAss(assRepository.findByAssid(interViewUser.getInterViewAss().getAssId()));
            } else if (wxuser.getSecondAss() == null) {
                wxuser.setSecondAss(assRepository.findByAssid(interViewUser.getInterViewAss().getAssId()));
            }
            wxuserRepository.save(wxuser);
            jsonObject.put("status", 1);
            resultModel.setData(jsonObject);
            resultModel.setMsg("成功加入:" + interViewUser.getInterViewAss().getAssName() + "!");
            return resultModel;
        }
        return null;
    }

    /**
     * 学生拒绝加入社团
     *
     * @param id
     * @return
     */
    @GetMapping("wxapi/interview/studentRefuse")
    @ResponseBody
    public ResultModel studentRefuse(@RequestParam("id") String id) {
        InterViewUser interViewUser = interViewUserRepository.getOne(id);
        InterViewEnum interViewEnum = interViewUser.getInterViewStatus();
        if (interViewEnum == InterViewEnum.Stage_Two_Success) {
            interViewEnum = InterViewEnum.Stage_Three_Failed;
            //更新状态
            interViewUser.setInterViewStatus(interViewEnum);
            //设置控制按钮
            interViewUser.setButtonControl(0);
            interViewUserRepository.save(interViewUser);
        }
        return new ResultModel(20000, "");
    }

    /**
     * 获取学生的面试记录
     *
     * @param studentId
     * @return
     */
    @GetMapping("wxapi/interview/getUserInterviews")
    @ResponseBody
    public ResultModelArrary getUserInterviews(@RequestParam("studentId") String studentId) {
        if (studentId.length()==0) {
            return new ResultModelArrary(20000, "查找失败，请重新进入系统");
        }
        ArrayList<InterViewUser> interViewUsers = interViewUserRepository.findAllByStudentId(studentId);
        net.sf.json.JSONArray jsonArray = new net.sf.json.JSONArray();
        for (InterViewUser interViewUser : interViewUsers) {
            JSONObject jsonObject = interViewUser.getInfo();
            if (interViewUser.getInterViewStatus().getStep() != 3) {
                jsonObject.put("backMessage", interViewUser.getInterViewAss().getShowMessage());
            } else
                jsonObject.put("backMessage", interViewUser.getInterViewAss().getConfirmJoinMessage());
            jsonArray.add(jsonObject);
        }
        return new ResultModelArrary(20000, "", jsonArray);
    }

    @GetMapping("wxapi/interview/goGet/{api}")
    @ResponseBody
    public JSONObject goApiGet(HttpServletRequest httpServletRequest, @PathVariable("api") String api) {
        String queryString = httpServletRequest.getQueryString();
        try {
            String encode = URLDecoder.decode(queryString, "UTF-8");
            String url = "http://localhost:8888/" + api + "?" + encode;
            RestTemplate restTemplate = new RestTemplate();
            String resp = restTemplate.getForEntity(url, String.class).getBody();
            return JSONObject.parseObject(resp);
        } catch (UnsupportedEncodingException e) {
            e.printStackTrace();
        }
        return null;
    }

    @PostMapping("wxapi/interview/goPost/{api}")
    @ResponseBody
    public JSONObject goApiPost(HttpServletRequest httpServletRequest,
                                @PathVariable("api") String api,
                                @RequestParam MultiValueMap<String, Object> map) {
        String url = "http://localhost:8888/" + api;
        //获取请求中的请求题
        BufferedReader br = null;
        try {
            br = new BufferedReader(new InputStreamReader(httpServletRequest.getInputStream(), "UTF-8"));
        } catch (IOException e) {
            e.printStackTrace();
        }
        String line = null;
        StringBuilder sb = new StringBuilder();
        try {
            while ((line = br.readLine()) != null) {
                sb.append(line);
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        //发送post请求
        RestTemplate restTemplate = new RestTemplate();
        HttpHeaders headers = new HttpHeaders();
        //设置请求头为json类型
        MediaType type = null;
        String s = null;
        type = MediaType.parseMediaType("application/json; charset=UTF-8");
        headers.setContentType(type);
        headers.add("Accept", MediaType.APPLICATION_JSON.toString());
        //拼接请求体和请求头
        HttpEntity<String> formEntity = new HttpEntity<String>(sb.toString(), headers);
        s = restTemplate.postForEntity(url, formEntity, String.class).getBody();

        return JSONObject.parseObject(s);
    }

    @PostMapping("wxapi/interview/goPostFormData/{api}")
    @ResponseBody
    public JSONObject goApiPostFormData(@PathVariable("api") String api,
                                        @RequestParam("file") MultipartFile file,
                                        @RequestParam("department") String department) throws IOException {
        String url = "http://localhost:8888/" + api;
        String tempFileName = "/tmp/" + file.getOriginalFilename();
        FileOutputStream fo = new FileOutputStream(tempFileName);
        fo.write(file.getBytes());
        fo.close();
        MultiValueMap<String, Object> map = new LinkedMultiValueMap<>();
        map.add("department", department);
        map.add("file", new FileSystemResource(tempFileName));

        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.MULTIPART_FORM_DATA);
        headers.add("Accept", MediaType.APPLICATION_JSON.toString());
        HttpEntity<MultiValueMap<String, Object>> formEntity = new HttpEntity<MultiValueMap<String, Object>>(map, headers);
        RestTemplate restTemplate = new RestTemplate();
        String s = restTemplate.postForEntity(url, formEntity, String.class).getBody();

        return JSONObject.parseObject(s);
    }
}
