package com.fran6k.study.springbootstudy.bean.WxBean;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.fran6k.study.springbootstudy.bean.Ass;
import lombok.Data;
import org.hibernate.annotations.GenericGenerator;

import javax.persistence.*;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

@Entity
@Table(name = "Wxuser")
@Data
public class Wxuser {
    @Id
    @GeneratedValue(generator = "paymentableGenerator")
    @GenericGenerator(name = "paymentableGenerator", strategy = "uuid")
    private String id;
    private String nickname;
    private String openId;
    private String realName;
    private String avatar;
    private String studentId;
    private String weChatId;
    private String phoneNum;


    @ManyToOne(targetEntity= Ass.class)
    @JoinColumn(name="assEntityFirst",referencedColumnName = "assid")
    private Ass firstAss;

    @ManyToOne(targetEntity= Ass.class)
    @JoinColumn(name="assEntitySecond",referencedColumnName = "assid")
    private Ass secondAss;

    public Wxuser(){
        this.nickname=null;
        this.openId=null;
        this.realName=null;
        this.avatar=null;
    }


    public List<String> getAssName(){
        List<String> name = new ArrayList<>();
        try {
            name.add(firstAss.getAssname());
        } catch (NullPointerException e){}
        try {
            name.add(secondAss.getAssname());
        } catch (NullPointerException e){}
        return name;
    }
    public JSONArray getAssObject() {
        JSONObject assFirst = new JSONObject();
        String assName = "-1";
        String assLogo = "-1";
        try {
            assName = firstAss.getAssname();
            assLogo = firstAss.getLogo();
        } catch (NullPointerException e) {} finally {
            assFirst.put("assName",assName);
            assFirst.put("assLogo",assLogo);
        }
        JSONObject assSec = new JSONObject();
        assName = "-1";
        assLogo = "-1";
        try {
            assName = secondAss.getAssname();
            assLogo = secondAss.getLogo();
        } catch (NullPointerException e) {} finally {
            assSec.put("assName",assName);
            assSec.put("assLogo",assLogo);
        }
        JSONArray jsonArray = new JSONArray();
        jsonArray.add(assFirst);
        jsonArray.add(assSec);
//        System.out.println(jsonArray);
        return jsonArray;
    }
    public JSONObject getInfo(){
        JSONObject jsonObject = new JSONObject();
        jsonObject.put("id",id);
        jsonObject.put("nickName",nickname);
        jsonObject.put("openId",openId);
        jsonObject.put("realName",realName);
        jsonObject.put("avatar",avatar);
        jsonObject.put("studentId",studentId);
        jsonObject.put("weChatId",weChatId);
        jsonObject.put("phoneNum",phoneNum);
        List<String> list = this.getAssName();
        jsonObject.put("assList",list);
        return jsonObject;
    }

    public String getName(){
        return nickname;
    }
}
