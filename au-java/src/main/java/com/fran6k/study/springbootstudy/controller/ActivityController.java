package com.fran6k.study.springbootstudy.controller;

import com.alibaba.fastjson.JSONObject;
import com.fran6k.study.springbootstudy.bean.InReviewBean.Activity;
import com.fran6k.study.springbootstudy.bean.InReviewBean.ActivityComplete;
import com.fran6k.study.springbootstudy.dao.ActivityCompleteRepository;
import com.fran6k.study.springbootstudy.dao.ActivityRepository;
import com.fran6k.study.springbootstudy.dao.AssRepository;
import com.fran6k.study.springbootstudy.utils.ResultModelArrary;
import com.fran6k.study.springbootstudy.utils.UUIDUtils;
import net.sf.json.JSONArray;
import org.apache.tomcat.util.http.fileupload.IOUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import javax.servlet.http.HttpServletResponse;
import java.io.*;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
@Controller
@CrossOrigin

/**
 *  活动提交请求
 */
public class ActivityController {
    /*
    6.审核完成
    5.等待老师审核
    4.等待财务部工作人员审核
    3.等待提交财务报表
    2.等待策划部工作人员审核
     */
    public static int finalstep = 6;

    @Autowired
    private ActivityRepository activityRepository;

    @Autowired
    private ActivityCompleteRepository activityCompleteRepository;

    @Autowired
    private AssRepository assRepository;


    /**
     * 提交社团活动
     * @param files 文件
     * @param description 活动简介
     * @param date 活动日期
     * @param assid 社团ID
     * @param name 活动名称
     * @param filed 文件类型，1策划书，2财务报表
     * @param step 当前进度
     * @return JSON数据
     */
    @PostMapping("activity/submit")
    @ResponseBody
    public JSONObject submit(@RequestParam("files") MultipartFile files,
                             @RequestParam(value = "description",required = false) String description,
                             @RequestParam(value = "date",required = false) Date date,
                             @RequestParam(value = "name",required = false) String name,
                             @RequestParam("filed") int filed,
                             @RequestParam("assid") int assid,
                             @RequestParam("step") int step){
        Activity activity = activityRepository.findByAssidAndStepIsNot(assid,finalstep);

        if(activity==null){
            activity = new Activity();
        }
        else {
            if ((activity.getFile1() != null)&& filed==1) {
                File deletefile = new File(activity.getFile1());
                deletefile.delete();

                System.out.println("删除成功");
            }
            if ((activity.getFile2() != null)&& filed==2) {
                File deletefile = new File(activity.getFile2());
                deletefile.delete();
                System.out.println("删除成功");

            }
        }
        if(filed==1){
            activity.setFile1(this.upload(files));
            activity.setName(name);
            activity.setDate(date);
            activity.setDescription(description);
        }
        if(filed==2)activity.setFile2(this.upload(files));

        activity.setAssid(assid);

        activity.setStep(++step);
        activity.setBackmsg(null);
        activityRepository.save(activity);

        JSONObject result =new JSONObject();
        result.put("code",20000);
        result.put("message","ok");
        return result;

    }
    public String upload(MultipartFile file){

//        System.out.println(file);
        File fileDir = new File(rootPath);
        if (!fileDir.exists() && !fileDir.isDirectory())
            fileDir.mkdirs();

        //使用uuid工具
        String uuid = UUIDUtils.getUUID()+"";
        //创建日期
        SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd");
        Date date = new Date();
        String dateStr = simpleDateFormat.format(date);
        String storagePath = rootPath +"/"+uuid+"_"+file.getOriginalFilename();

        try {
            file.transferTo(new File(storagePath));
        } catch (IOException e) {
            e.printStackTrace();
        }
        return storagePath;
    }

//    当社长进入页面的时候，获取当前活动进度和失败信息
    @GetMapping("activity/getinfo")
    @ResponseBody
    public ResultModelArrary getinfo(@RequestParam("assid") int assid){
        Activity activity = activityRepository.findByAssidAndStepIsNot(assid,finalstep);
        ResultModelArrary resultModelArrary = new ResultModelArrary(20000,"ok");
        if(activity!=null){
            resultModelArrary.setData(JSONArray.fromObject(activity));
        }
        return resultModelArrary;
    }

    //审核时候，审核成功的下一步
    @GetMapping("activity/nextstep")
    @ResponseBody
    public JSONObject nextstep(@RequestParam("assid") int assid){
        Activity activity = activityRepository.findByAssidAndStepIsNot(assid,finalstep);
        int temstep = activity.getStep();
        activity.setStep(++temstep);
        activityRepository.save(activity);

        if(activity.getStep()==finalstep){
            ActivityComplete activityComplete = new ActivityComplete();
            activityComplete.setActivityName(activity.getName());
            activityComplete.setTwitter("");
            activityComplete.setDate(activity.getDate());
            activityComplete.setFile1(activity.getFile1());
            activityComplete.setFile2(activity.getFile2());
            activityComplete.setDescription(activity.getDescription());
            activityComplete.setAss(assRepository.findByAssid(activity.getAssid()));
            activityCompleteRepository.save(activityComplete);
        }


        JSONObject result =new JSONObject();
        result.put("code",20000);
        result.put("message","ok");
        return result;
    }

    //审核失败的时候设置返回信息
    @GetMapping("activity/setbackmsg")
    @ResponseBody
    public ResultModelArrary setbackmsg(@RequestParam("assid") int assid,
                                        @RequestParam("backmsg") String _backmsg){

        Activity activity = activityRepository.findByAssidAndStepIsNot(assid,finalstep);
        int temstep = activity.getStep();
        activity.setStep(1);
        activity.setBackmsg(_backmsg);
        activityRepository.save(activity);
        ResultModelArrary resultModelArrary = new ResultModelArrary(20000,"ok");
        return resultModelArrary;
    }

    //获取审核列表
    @GetMapping("activity/getactivities")
    @ResponseBody
    public ResultModelArrary getactivities(){
        ArrayList<Activity> activities = activityRepository.findAllByStepNot(finalstep);
        ResultModelArrary resultModelArrary = new ResultModelArrary(20000,"ok",JSONArray.fromObject(activities));
        return resultModelArrary;
    }

    //文件上传
    @Value("${upload.activity.path}")
    private String rootPath;


    @GetMapping("activity/downFile")
    @ResponseBody
    public void downFile(@RequestParam("path") String path, HttpServletResponse response) throws Exception{
        File file = new File(path);
        String[] filename = file.getName().split("_");
        FileInputStream fileInputStream = null;
        fileInputStream = new FileInputStream(file);
        response.setContentType("application/force-download");
        response.setCharacterEncoding("UTF-8");
        response.setHeader("Content-Disposition", "attachment;fileName=" +   java.net.URLEncoder.encode(filename[1],"UTF-8"));
        IOUtils.copy(fileInputStream ,response.getOutputStream());
    }

}
