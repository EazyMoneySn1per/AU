package com.fran6k.study.springbootstudy.controller;

import com.alibaba.fastjson.JSONObject;
import com.fran6k.study.springbootstudy.bean.InReviewBean.Activity;
import com.fran6k.study.springbootstudy.bean.InReviewBean.ActivityComplete;
import com.fran6k.study.springbootstudy.bean.Outlay;
import com.fran6k.study.springbootstudy.bean.Twitter;
import com.fran6k.study.springbootstudy.dao.ActivityCompleteRepository;
import com.fran6k.study.springbootstudy.dao.ActivityRepository;
import com.fran6k.study.springbootstudy.dao.OutlayRepository;
import com.fran6k.study.springbootstudy.dao.TwitterRepository;
import com.fran6k.study.springbootstudy.utils.ResultModel;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.testng.annotations.Test;

import java.util.ArrayList;
import java.util.Calendar;
import java.util.Date;

/**
 *  数据可视化接口
 * @author frankkkkkk
 */
@Controller
public class DataShowController {
    @Autowired
    ActivityRepository activityRepository;
    @Autowired
    ActivityCompleteRepository activityCompleteRepository;
    @Autowired
    TwitterRepository twitterRepository;
    @Autowired
    OutlayRepository outlayRepository;

    /**
     * 获取正在审核中的活动财务报表推文
     * @return jsonObject
     */
    @GetMapping("dataShow/getUnfinishedEvent")
    @ResponseBody
    public ResultModel getUnfinishedEvent(){
        ArrayList<Activity> activities= activityRepository.findAllByStepNot(ActivityController.finalstep);
        ArrayList<Twitter> twitters = twitterRepository.findAllByStepNot(TwitterController.finalstep);
        ArrayList<Outlay> outlays = outlayRepository.findAllByStepNot(OutlayController.finalstep);
        int activitiesNum = activities.size();
        int twittersNum = twitters.size();
        int outlaysNum = outlays.size();
        JSONObject jsonObject = new JSONObject();
        jsonObject.put("activitiesNum",activitiesNum);
        jsonObject.put("twittersNum",twittersNum);
        jsonObject.put("outlaysNum",outlaysNum);
        return new ResultModel(20000,"",jsonObject);
    }
    /**
     * 获取所有活动财务报表推文
     * @return ResultModel
     */
    @GetMapping("dataShow/getFinishedEvent")
    @ResponseBody
    public ResultModel getFinishedEvent(){
        ArrayList<ActivityComplete> activities= activityCompleteRepository.findAll();
        ArrayList<Twitter> twitters = twitterRepository.findAllByStep(TwitterController.finalstep);
        ArrayList<Outlay> outlays = outlayRepository.findAllByStep(OutlayController.finalstep);
        int activitiesNum = activities.size();
        int twittersNum = twitters.size();
        int outlaysNum = outlays.size();
        JSONObject jsonObject = new JSONObject();
        jsonObject.put("activitiesNum",activitiesNum);
        jsonObject.put("twittersNum",twittersNum);
        jsonObject.put("outlaysNum",outlaysNum);
        return new ResultModel(20000,"",jsonObject);
    }
    /**
     * 获取今日提交的活动，财务报表，推文
     * @return ResultModel
     */
    @GetMapping("dataShow/getTodayEvent")
    @ResponseBody
    public ResultModel getTodayEvent(){
        ArrayList<Activity> activities = activityRepository.findTodayAll();
        ArrayList<Twitter> twitters = twitterRepository.findTodayAll();
        ArrayList<Outlay> outlays = outlayRepository.findTodayAll();

        int twittersNum = twitters.size();
        int activitiesNum = activities.size();
        int outlaysNum = outlays.size();
        JSONObject jsonObject = new JSONObject();
        jsonObject.put("activitiesNum",activitiesNum);
        jsonObject.put("twittersNum",twittersNum);
        jsonObject.put("outlaysNum",outlaysNum);
        return new ResultModel(20000,"",jsonObject);
    }
}
