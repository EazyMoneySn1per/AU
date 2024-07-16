package com.fran6k.study.springbootstudy.utils;

import com.alibaba.fastjson.JSONObject;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.converter.StringHttpMessageConverter;
import org.springframework.web.client.RestTemplate;

import java.nio.charset.Charset;
import java.nio.charset.StandardCharsets;

public class UIMLoginUtil {
//    @Value("${UIMKey}")
    private static String key="Rh&z83I7X7G0vCLf";

    static public JSONObject veri(String username,String password){
        String url = "http://10.1.20.86/api/tripartite/auth";
        RestTemplate restTemplate = new RestTemplate();
        //设置编码类型
        restTemplate.getMessageConverters().set(1,new StringHttpMessageConverter(StandardCharsets.UTF_8));
        //加密的数据
        JSONObject jsonObject = new JSONObject();
        jsonObject.put("username",username);
        jsonObject.put("password",password);
        //进行AES加密
        String encryptData = AESUtil.encrypt(jsonObject.toString(),key);
        HttpHeaders headers = new HttpHeaders();
        //设置请求类型
        headers.setContentType(MediaType.TEXT_PLAIN);
        //组装请求头和请求体
        HttpEntity<String>formEntity = new HttpEntity<>(encryptData,headers);
        //String.class返回结果的映射类型
        String res = restTemplate.postForObject(url,formEntity,String.class);
        //AES解密
        System.out.println("私钥"+key);
        System.out.println("解密前的消息:" + res);
        String decryptData = AESUtil.decrypt(res,key);
        System.out.println("解密后的消息:" + decryptData);
        JSONObject jsonObject1 = JSONObject.parseObject(decryptData);
        System.out.println("转换成JSON对象" + jsonObject1);
        return jsonObject1;
    }
}
