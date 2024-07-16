package com.fran6k.study.springbootstudy.bean;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.fran6k.study.springbootstudy.bean.WxBean.Wxuser;
import com.fran6k.study.springbootstudy.dao.AssRepository;
import com.fran6k.study.springbootstudy.dao.UserRepository;
import com.fran6k.study.springbootstudy.dao.WxuserRepository;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

@RunWith(SpringRunner.class)
@SpringBootTest

public class AssTestTwo {
    @Autowired
    private AssRepository assRepository;
    @Autowired
    private WxuserRepository wxuserRepository;
    @Test
    public void test(){
//        List<Ass> ass = assRepository.findAll();
//        for(Ass s : ass){
//            System.out.println(s.getAssname());
//            List<User> users= s.getList();
//            System.out.println(users);
////            for(User u:users){
////                System.out.println(u);
////            }
//        }
//        Wxuser wxuser = wxuserRepository.findByOpenId("ofM8f59yb_AkxsUdYvjoin1YNByo");
//        System.out.println(wxuser.getAssName());
//        System.out.println(wxuser.getAssName());
//        user.setAss(ass);
//        userRepository.save(user);
//        System.out.println(ass);
//        System.out.println(user);
        Ass ass = assRepository.findByAssname("音乐社");
        System.out.println(ass.toString());

    }
}