package com.fran6k.study.springbootstudy.controller;

import com.alibaba.fastjson.JSONObject;
import com.fran6k.study.springbootstudy.bean.InReviewBean.ActivityComplete;
import com.fran6k.study.springbootstudy.bean.Ass;
import com.fran6k.study.springbootstudy.bean.Outlay;
import com.fran6k.study.springbootstudy.bean.Twitter;
import com.fran6k.study.springbootstudy.bean.WxBean.InterViewAss;
import com.fran6k.study.springbootstudy.bean.WxBean.Wxuser;
import com.fran6k.study.springbootstudy.dao.*;
import com.fran6k.study.springbootstudy.utils.ResultModel;
import com.fran6k.study.springbootstudy.utils.ResultModelArrary;
import net.sf.json.JSONArray;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;

import java.io.IOException;
import java.util.List;
import java.util.Map;

import static com.fran6k.study.springbootstudy.controller.TwitterController.finalstep;

@Controller
public class AdminController {
    @Autowired
    AssRepository assRepository;
    @Autowired
    WxuserRepository wxuserRepository;
    @Autowired
    ActivityCompleteRepository activityCompleteRepository;
    @Autowired
    TwitterRepository twitterRepository;
    @Autowired
    OutlayRepository outlayRepository;
    @Autowired
    InterViewAssRepository interViewAssRepository;

    /**
     * 获取所有社团
     * @return
     * @throws IOException
     */
    @GetMapping("admin/getAssociations")
    @ResponseBody
    public ResultModelArrary getAssociation() throws IOException {
        JSONArray jsonArray = new JSONArray();
        List<Ass> associations= assRepository.findAll();
        for(Ass s : associations){
            jsonArray.add(s.getAss());
        }
        ResultModelArrary resultModel = new ResultModelArrary(20000, "",jsonArray);
        return resultModel;
    }
    @GetMapping("admin/getAssociationsByAssid")
    @ResponseBody
    public ResultModel getAssociationsByAssid(@RequestParam("assid")int assId) throws IOException {
        Ass ass = assRepository.findByAssid(assId);
        JSONObject jsonObject = ass.getAss();
        return new ResultModel(20000, "",jsonObject);
    }

    /**
     * 更改社团信息
     * @param map
     * @return
     * @throws IOException
     */
    @PostMapping ("admin/setAssociationsByAssid")
    @ResponseBody
    public ResultModel setAssociationsByAssid(@RequestBody Map map) throws IOException {
        int assId = (int) map.get("assId");
        String assName = map.get("assName").toString();
        String assType = map.get("assType").toString();
        double money = Double.parseDouble(map.get("money").toString());
        int isExecllent = (int) map.get("isExecllent");
        String presidentName = map.get("presidentName").toString();
        String presidentId = map.get("presidentId").toString();
        String teacherName = map.get("teacherName").toString();
        String teacherPhone = map.get("teacherPhone").toString();
        String teacherPost = map.get("teacherPost").toString();
        String assDescription = map.get("assDescription").toString();


        Ass ass = assRepository.findByAssid(assId);
        if (!ass.getAssname().equals(assName)) ass.setAssname(assName);
        if (!ass.getAsstype().equals(assType)) ass.setAsstype(assType);
        ass.setMoney(money);
        ass.setIsExecllent(isExecllent);
        if (!ass.getPresidentname().equals(presidentName)) ass.setPresidentname(presidentName);
        if (!ass.getPresidentId().equals(presidentId)) ass.setPresidentId(presidentId);
        if (!ass.getTeachername().equals(teacherName)) ass.setTeachername(teacherName);
        if (!ass.getTeacherphone().equals(teacherPhone)) ass.setTeacherphone(teacherPhone);
        if (!ass.getTeacherpost().equals(teacherPost)) ass.setTeacherpost(teacherPost);
        if (!ass.getAssDescription().equals(assDescription)) ass.setAssDescription(assDescription);

        assRepository.save(ass);
        return new ResultModel(20000, "");
    }

    /**
     * 新增社团
     * @param map
     * @return
     * @throws IOException
     */
    @PostMapping ("admin/addAssociations")
    @ResponseBody
    public ResultModel addAssociations(@RequestBody Map map) throws IOException {
        String assName = map.get("assName").toString();
        String assType = map.get("assType").toString();
        double money = Double.parseDouble(map.get("money").toString());
        int isExecllent = (int) map.get("isExecllent");
        String presidentName = map.get("presidentName").toString();
        String presidentId = map.get("presidentId").toString();
        String teacherName = map.get("teacherName").toString();
        String teacherPhone = map.get("teacherPhone").toString();
        String teacherPost = map.get("teacherPost").toString();
        String assDescription = map.get("assDescription").toString();

        Ass ass = new Ass();
        ass.setAssname(assName);
        ass.setAsstype(assType);
        ass.setMoney(money);
        ass.setIsExecllent(isExecllent);
        ass.setPresidentname(presidentName);
        ass.setPresidentId(presidentId);
        ass.setTeachername(teacherName);
        ass.setTeacherphone(teacherPhone);
        ass.setTeacherpost(teacherPost);
        ass.setAssDescription(assDescription);
        ass.setLogo("default.png");
        assRepository.save(ass);

        // 同步建立社团面试表中的部门
        InterViewAss interViewAss = new InterViewAss();
        interViewAss.setAssId(ass.getAssid());
        interViewAss.setAssName(assName);
        interViewAss.setLogo("default.png");
        interViewAss.setPresidentName("无");
        interViewAss.setPresidentWechat("无");
        interViewAss.setShowMessage("无");
        interViewAss.setConfirmJoinMessage("无");
        interViewAss.setCodeUrl("无");
        interViewAssRepository.save(interViewAss);

        return new ResultModel(20000, "添加成功");
    }

    @GetMapping("admin/deleteAssociations")
    @ResponseBody
    public ResultModel deleteAssociations(@RequestParam("assid")int assId) throws IOException {
        Ass ass = assRepository.findByAssid(assId);
        if (ass != null) assRepository.delete(ass);
        else return new ResultModel(20000, "社团不存在");
        return new ResultModel(20000, "删除成功");
    }

    @GetMapping("admin/getAssociationsNameMapAssid")
    @ResponseBody
    public ResultModel getAssociationsNameMapAssid() throws IOException {
        JSONObject jsonObject = new JSONObject();
        List<Ass> associations= assRepository.findAll();
        for(Ass s : associations){
            jsonObject.put(String.valueOf(s.getAssid()),s.getAssname());
        }
        ResultModel resultModel = new ResultModel(20000, "",jsonObject);
        return resultModel;
    }

    @GetMapping("admin/getAssociationsName")
    @ResponseBody
    public ResultModelArrary getAssociationName() throws IOException {
        JSONArray jsonArray = new JSONArray();
        List<Ass> associations= assRepository.findAll();
        for(Ass s : associations){
            jsonArray.add(s.getAssname());
        }
        ResultModelArrary resultModel = new ResultModelArrary(20000, "",jsonArray);
        return resultModel;
    }

    @GetMapping("admin/getAssociationsStudents")
    @ResponseBody
    public ResultModelArrary getAssociationStudents(@RequestParam("assId") int assId) {
        JSONArray jsonArray = new JSONArray();

        Ass ass = assRepository.findByAssid(assId);
        List<Wxuser> list = wxuserRepository.findAllByFirstAssOrSecondAss(ass,ass);
        for(Wxuser wxuser:list){
            jsonArray.add(wxuser.getInfo());
        }
        ResultModelArrary resultModel = new ResultModelArrary(20000, "",jsonArray);
        return resultModel;
    }
    //根据条件筛选查询学生
    @GetMapping("admin/getStudentsList")
    @ResponseBody
    public ResultModel getStudentsList(@RequestParam(value = "page",required = true) int page,
                                       @RequestParam(value = "limit",required = true) int limit,
                                       @RequestParam(value = "nickName",required = false) String nickName,
                                       @RequestParam(value = "realName",required = false) String realName,
                                       @RequestParam(value = "studentId",required = false) String studentId,
                                       @RequestParam(value = "weChatId",required = false) String weChatId,
                                       @RequestParam(value = "phoneNum",required = false) String phoneNum,
                                       @RequestParam(value = "assName",required = false) String assName,
                                       @RequestParam(value = "assId",required = false) String assId) {

        JSONArray jsonArray = new JSONArray();
        //查询的页数和条数
        PageRequest pageRequest = PageRequest.of(page-1,limit);
        //判断有无assName，没有的话assid就是null
        if(assName != null && assId == null){
            Ass ass = assRepository.findByAssname(assName);
            assId = String.valueOf(ass.getAssid());
        }
        //根据条件查询
        Page<Wxuser> queryPage = wxuserRepository.searchAllBy(nickName,realName,studentId,weChatId,phoneNum,assId,pageRequest);
        List<Wxuser> list = queryPage.getContent();

        for(Wxuser wxuser:list){
            jsonArray.add(wxuser.getInfo());
        }
        JSONObject jsonObject = new JSONObject();
        jsonObject.put("item",jsonArray);
        jsonObject.put("total",queryPage.getTotalElements());
        ResultModel resultModel = new ResultModel(20000,"",jsonObject);
        return resultModel;
    }
    //根据条件筛选查询活动
    @GetMapping("admin/getActivityCompleteList")
    @ResponseBody
    public ResultModel getActivityCompleteList(@RequestParam(value = "page",required = true) int page,
                                       @RequestParam(value = "limit",required = true) int limit,
                                       @RequestParam(value = "assName",required = false) String assName,
                                       @RequestParam(value = "assId",required = false) String assId) {

//        JSONArray jsonArray = new JSONArray();
        //查询的页数和条数
        PageRequest pageRequest = PageRequest.of(page-1,limit);

        int assidInt = -1;
        if (assId != null)assidInt = Integer.parseInt(assId);

        Page<ActivityComplete> queryPage = null;
        if(assidInt != -1){
            queryPage = activityCompleteRepository.findAllByAss_AssnameOrAss_AssidOrderByDate(assName, assidInt,pageRequest);
        }
        else
            queryPage = activityCompleteRepository.findAll(pageRequest);
        List<ActivityComplete> list = queryPage.getContent();

        JSONArray jsonArray = new JSONArray();
        for(ActivityComplete e : list){
            jsonArray.add(e.getInfo());
        }
        JSONObject jsonObject = new JSONObject();
        jsonObject.put("item",jsonArray);
        jsonObject.put("total",queryPage.getTotalElements());
        ResultModel resultModel = new ResultModel(20000,"",jsonObject);
        return resultModel;
    }

    //根据条件筛选查询推文
    @GetMapping("admin/getTwitterList")
    @ResponseBody
    public ResultModel getTwitter(@RequestParam(value = "page",required = true) int page,
                                           @RequestParam(value = "limit",required = true) int limit,
                                           @RequestParam(value = "assName",required = false) String assName,
                                           @RequestParam(value = "assId",required = false) String assId){
        PageRequest pageRequest = PageRequest.of(page-1,limit);
        int assidInt = -1;
        if (assId != null)assidInt = Integer.parseInt(assId);
        Page<Twitter> twitters;
        if(assidInt != -1){
             twitters = twitterRepository.findAllByAssidAndStep(assidInt,finalstep,pageRequest);
        }
        else{
            twitters = twitterRepository.findAllByStep(finalstep,pageRequest);
        }
        List<Twitter> list = twitters.getContent();

        JSONObject jsonObject = new JSONObject();
        jsonObject.put("item",JSONArray.fromObject(list));
        jsonObject.put("total",twitters.getTotalElements());
        ResultModel resultModel = new ResultModel(20000,"",jsonObject);
        return resultModel;
    }

    @PostMapping("admin/setActivityTwitter")
    @ResponseBody
//    public ResultModel setActivityTwitter(@RequestParam("id") int id,
//                                          @RequestParam("twitterUrl") String twitterUrl){
    public ResultModel setActivityTwitter(@RequestBody Map map){
        String id = map.get("id").toString();
        String twitterUrl = map.get("twitterUrl").toString();
        ActivityComplete activityComplete = activityCompleteRepository.findById(Integer.parseInt(id));
        activityComplete.setTwitter(twitterUrl);
        System.out.println(activityComplete.getTwitter());
        activityCompleteRepository.save(activityComplete);
        return new ResultModel(20000,"");
    }
    //根据条件筛选查询经费报销情况
    @GetMapping("admin/getOutlayCompleteList")
    @ResponseBody
    public ResultModel getOutlayCompleteList(@RequestParam(value = "page",required = true) int page,
                                  @RequestParam(value = "limit",required = true) int limit,
                                  @RequestParam(value = "assName",required = false) String assName,
                                  @RequestParam(value = "assId",required = false) String assId){
        PageRequest pageRequest = PageRequest.of(page-1,limit);

        int assidInt = -1;
        if (assId != null)assidInt = Integer.parseInt(assId);
        Page<Outlay> outlays;
        if(assidInt != -1){
            outlays = outlayRepository.findAllByAssidAndStep(assidInt,finalstep,pageRequest);
        }
        else{
            outlays = outlayRepository.findAllByStep(finalstep,pageRequest);
        }
        List<Outlay> list = outlays.getContent();

        JSONObject jsonObject = new JSONObject();
        jsonObject.put("item",JSONArray.fromObject(list));
        jsonObject.put("total",outlays.getTotalElements());
        ResultModel resultModel = new ResultModel(20000,"",jsonObject);
        return resultModel;
    }

    @GetMapping("admin/test")
    @ResponseBody
    public String test(){
        String url = "https://mp.weixin.qq.com/s/ZClnmFzROEIgPUlkPqTAXw";
        RestTemplate restTemplate=new RestTemplate();
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_FORM_URLENCODED);

        ResponseEntity<String> responseEntity = restTemplate.getForEntity(url,String.class);
        return responseEntity.getBody();
    }
}
