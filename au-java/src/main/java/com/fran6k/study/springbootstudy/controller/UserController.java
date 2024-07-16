package com.fran6k.study.springbootstudy.controller;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.fran6k.study.springbootstudy.bean.User;
import com.fran6k.study.springbootstudy.bean.Vo.VoUser;
import com.fran6k.study.springbootstudy.bean.Vo.VoUserInfo;
import com.fran6k.study.springbootstudy.dao.UserRepository;
import com.fran6k.study.springbootstudy.utils.AESUtil;
import com.fran6k.study.springbootstudy.utils.JwtUtil;
import com.fran6k.study.springbootstudy.bean.Vo.VoToken;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import java.util.Arrays;
import java.util.List;

@Controller
public class UserController {
    @Autowired
    private UserRepository userRepository;

    @RequestMapping(value = "auth/login", method = RequestMethod.POST)
    @ResponseBody
//    @CrossOrigin
    public JSONObject getuser(@RequestBody String Info) throws Exception {
//        System.out.println(Info);
        // AES解密
        String decrypt = AESUtil.decrypt(Info, "RC3!%(a14f*op52E");
        JSONObject jsonObject = JSONObject.parseObject(decrypt);
        VoUser voUser = JSON.toJavaObject(jsonObject,VoUser.class);
        User user = userRepository.findByNameAndPassword(voUser.getUsername(), voUser.getPassword());
        JSONObject result = new JSONObject();
        if (user != null) {
            //Jwt签名，存入用户姓名
            String token = JwtUtil.sign(voUser.getUsername(), "-1");
            VoToken votoken = new VoToken();
            votoken.setToken(token);
            result.put("code", 20000);
            result.put("data", votoken);

        }
        return result;

    }

    @GetMapping("/auth/info")
    @ResponseBody
    public JSONObject getinfo(@RequestParam("token") String token) {
        JSONObject result = new JSONObject();
        System.out.println(token);

        //1.验证token的有效性和合法性
        if (JwtUtil.verify(token)) {
            String username = JwtUtil.getUserName(token);

            User user = userRepository.findByName(username);

            VoUserInfo info = new VoUserInfo();
            info.setAvatar("");
            info.setName(user.getName());
            info.setAssid(user.getass().getAssid());
            info.setIntroduction("测试用户");
            //获取用户的权限
            List<String> roles = Arrays.asList(user.getRole());
            info.setRoles(roles);

            //设置返回的json
            result.put("data", info);
            result.put("code", 20000);
            result.put("message", "ok");

            return result;
        } else {
            return null;
        }
    }

    @PostMapping("/auth/logout")
    @ResponseBody
    public JSONObject logout(@RequestHeader("token") String token) {
        JSONObject result = new JSONObject();
        if (JwtUtil.verify(token)) {
            result.put("Message", "logout success");
            result.put("data", "success");
            result.put("code", 20000);
            return result;
        } else return null;
    }
}

