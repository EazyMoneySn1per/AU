package com.fran6k.study.springbootstudy.controller;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.fran6k.study.springbootstudy.bean.Ass;
import com.fran6k.study.springbootstudy.bean.User;
import com.fran6k.study.springbootstudy.dao.AssRepository;
import com.fran6k.study.springbootstudy.dao.UserRepository;
import com.fran6k.study.springbootstudy.utils.ResultModel;
import com.fran6k.study.springbootstudy.utils.ResultModelArrary;
import net.sf.json.JSONArray;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;

import java.util.ArrayList;
import java.util.List;

/**
 * 账号管理接口
 */
@Controller
public class AccountController {
    @Autowired
    UserRepository userRepository;
    @Autowired
    AssRepository assRepository;

    /**
     * 查询所有账号
     * @param role 根据权限查询，非必须
     * @param page 查询页数，必须
     * @param limit 每页条数，必须
     * @return
     */
    @GetMapping("account/getLists")
    @ResponseBody
    public ResultModel getLists(@RequestParam(value="role",required = false)String role,
                                @RequestParam(value="page",required = true)int page,
                                @RequestParam(value="limit",required = true)int limit){
        PageRequest pageRequest = PageRequest.of(page-1,limit);
        Page<User> queryPage;

        if(role!=null){
            queryPage = userRepository.findAllByRole(role,pageRequest);
        }
        else{
            queryPage = userRepository.findAll(pageRequest);
        }
        List<User> users = queryPage.getContent();
        JSONArray jsonArray = new JSONArray();
        for(User user: users){
            jsonArray.add(user.getInfo());
        }
        JSONObject jsonObject = new JSONObject();
        jsonObject.put("item",jsonArray);
        jsonObject.put("total",queryPage.getTotalElements());
        return new ResultModel(20000,"",jsonObject);
    }

    /**
     * 更变账号信息
     * @param name 用户名
     * @param password 密码
     * @param assName 所在社团，如果不对则无法更新数据
     * @param role 权限
     * @return
     */
    @PostMapping("account/updateInfo")
    @ResponseBody
    public ResultModel updateInfo(@RequestParam(value = "name")String name,
                                        @RequestParam(value = "password")String password,
                                        @RequestParam(value = "assName")String assName,
                                        @RequestParam(value = "role")String role){
        ResultModel resultModel = new ResultModel(20000,"");

        User user = userRepository.findByName(name);
        if (user == null) {
            user = new User();
        }
        user.setName(name);
        user.setPassword(password);
        user.setRole(role);
        try {
            Ass ass = assRepository.findByAssname(assName);
            user.setAss(ass);
            user.setAssid(ass.getAssid());
            userRepository.save(user);
            resultModel.setMsg("设置成功");
        } catch (NullPointerException e){
            e.printStackTrace();
            resultModel.setMsg("社团不存在，设置失败");
        }
        return resultModel;
    }

    /**
     * 删除用户
     * @param name
     * @return
     */
    @GetMapping("account/deleteUser")
    @ResponseBody
    public ResultModel deleteUser(@RequestParam(value = "name")String name){
        ResultModel resultModel = new ResultModel(20000,"");
        User user = userRepository.findByName(name);
        userRepository.delete(user);
        return resultModel;
    }
}
