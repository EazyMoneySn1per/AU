package com.fran6k.study.springbootstudy.controller;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.fran6k.study.springbootstudy.bean.WxBean.InterViewAss;
import com.fran6k.study.springbootstudy.bean.WxBean.InterViewEnum;
import com.fran6k.study.springbootstudy.bean.WxBean.InterViewUser;
import com.fran6k.study.springbootstudy.dao.InterViewAssRepository;
import com.fran6k.study.springbootstudy.dao.InterViewUserRepository;
import com.fran6k.study.springbootstudy.utils.ResultModel;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import java.util.List;
import java.util.Optional;

@Controller
public class InterviewController {
    @Autowired
    InterViewAssRepository interViewAssRepository;

    @Autowired
    InterViewUserRepository interViewUserRepository;
    /**
     * 获取社团的面试学生
     * @param assId
     * @param step
     * @return
     */
    @GetMapping("/interview/getAssInterviews")
    @ResponseBody
    public ResultModel getUserInterviews(@RequestParam(value="page",required = true)int page,
                                               @RequestParam(value="limit",required = true)int limit,
                                               @RequestParam("assId")String assId,
                                               @RequestParam(value = "interviewStep")int step){
        InterViewEnum interViewEnum = InterViewEnum.getByStep(step);
        InterViewAss interViewAss = interViewAssRepository.findByAssId(Integer.parseInt(assId));
        List<InterViewUser> interViewUsers =  interViewAss.getInterViewUsers();
        //获取满足面试状态的记录
        JSONArray jsonArray = new JSONArray();
        for(InterViewUser interViewUser:interViewUsers){
            if(interViewUser.getInterViewStatus() == interViewEnum){
                jsonArray.add(interViewUser.getInfo());
            }
        }
        //手动计算页数
        int firstIndex = (page-1)*limit;
        //取子串
        int lastIndex = firstIndex + limit;
        //边界判断
        if (lastIndex > jsonArray.size()) lastIndex = jsonArray.size();
        JSONArray resArray = new JSONArray();
        resArray.addAll(jsonArray.subList(firstIndex,lastIndex));
        JSONObject jsonObject = new JSONObject();
        jsonObject.put("item",resArray);
        jsonObject.put("total",jsonArray.size());
        jsonObject.put("totalInterviewUsers",interViewUsers.size());
        return new ResultModel(20000,"",jsonObject);
    }
    /**
     * 社长审核通过
     * @param id
     * @return
     */
    @GetMapping("/interview/AssConfirm")
    @ResponseBody
    public ResultModel AssConfirm(@RequestParam("id")String id){
        Optional<InterViewUser> optionalInterViewUser = interViewUserRepository.findById(id);
        InterViewUser interViewUser = optionalInterViewUser.get();
        InterViewEnum interViewEnum = interViewUser.getInterViewStatus();
        if (interViewEnum == InterViewEnum.Stage_One_WaitB){
            interViewEnum = InterViewEnum.Stage_Two_Success;
            //更新状态
            interViewUser.setInterViewStatus(interViewEnum);
            //更新控制按钮，使用权交给学生操作
            interViewUser.setButtonControl(2);
            interViewUserRepository.save(interViewUser);
        }
        return new ResultModel(20000,"");
    }

    /**
     * 社长拒绝
     * @param id
     * @return
     */
    @GetMapping("/interview/AssRefuse")
    @ResponseBody
    public ResultModel AssRefuse(@RequestParam("id")String id){
        InterViewUser interViewUser = interViewUserRepository.getOne(id);
        InterViewEnum interViewEnum = interViewUser.getInterViewStatus();
        if (interViewEnum == InterViewEnum.Stage_One_WaitB){
            interViewEnum = InterViewEnum.Stage_Two_Failed;
            //更新状态
            interViewUser.setInterViewStatus(interViewEnum);
            //设置控制按钮
            interViewUser.setButtonControl(0);
            interViewUserRepository.save(interViewUser);
        }
        return new ResultModel(20000,"");
    }
    @GetMapping("/interview/setAssMessage")
    @ResponseBody
    public ResultModel setAssShowMessage(@RequestParam("assId")String assId,
                                         @RequestParam("showMessage")String showMessage,
                                         @RequestParam("confirmJoinMessage")String confirmJoinMessage,
                                         @RequestParam("presidentName")String presidentName,
                                         @RequestParam("presidentWechat")String presidentWechat){
        InterViewAss interViewAss = interViewAssRepository.findByAssId(Integer.parseInt(assId));
        interViewAss.setShowMessage(showMessage);
        interViewAss.setConfirmJoinMessage(confirmJoinMessage);
        interViewAss.setPresidentWechat(presidentWechat);
        interViewAss.setPresidentName(presidentName);
        interViewAssRepository.save(interViewAss);
        return new ResultModel(20000,"");
    }
    @GetMapping("/interview/getAssMessage")
    @ResponseBody
    public ResultModel getAssMessage(@RequestParam("assId")String assId){
        InterViewAss interViewAss = interViewAssRepository.findByAssId(Integer.parseInt(assId));
        JSONObject jsonObject = new JSONObject();
        jsonObject.put("showMessage",interViewAss.getShowMessage());
        jsonObject.put("confirmJoinMessage",interViewAss.getConfirmJoinMessage());
        jsonObject.put("presidentWechat",interViewAss.getPresidentWechat());
        jsonObject.put("presidentName",interViewAss.getPresidentName());
        return new ResultModel(20000,"",jsonObject);
    }
}
