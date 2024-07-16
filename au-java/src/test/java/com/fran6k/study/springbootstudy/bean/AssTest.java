package com.fran6k.study.springbootstudy.bean;

import com.fran6k.study.springbootstudy.bean.WxBean.Wxuser;
import com.fran6k.study.springbootstudy.dao.AssRepository;
import com.fran6k.study.springbootstudy.dao.UserRepository;
import com.fran6k.study.springbootstudy.dao.WxuserRepository;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringRunner;

import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

@RunWith(SpringRunner.class)
@SpringBootTest

public class AssTest {
    @Autowired
    private AssRepository assRepository;
    @Autowired
    private UserRepository userRepository;
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
        User user =userRepository.findByName("admin");
        Ass ass = assRepository.findByAssid(3);
        for (int i=0;i<1000;i++) {
//            Wxuser wxuser = wxuserRepository.findByStudentId(String.valueOf(i));
            Wxuser wxuser = new Wxuser();
            wxuser.setStudentId(String.valueOf(10000+i));
            wxuser.setAvatar("1");
            wxuser.setNickname("1");
            wxuser.setRealName("测试学生"+(i+10000));
            wxuser.setWeChatId("1");
            wxuser.setOpenId("1");
            wxuser.setPhoneNum("1");
            wxuser.setFirstAss(ass);
//            wxuser.setRealName("测试学生"+(i+1000));
//            wxuser.setStudentId(String.valueOf(i));
            wxuserRepository.save(wxuser);
        }
//        ass.getUser().add(user);
//        user.setAss(ass);
//        userRepository.save(user);
//        System.out.println(ass);
//        System.out.println(user);
    }
}