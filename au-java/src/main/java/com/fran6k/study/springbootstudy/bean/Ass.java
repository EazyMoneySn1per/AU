package com.fran6k.study.springbootstudy.bean;

import com.alibaba.fastjson.JSONObject;
import com.fran6k.study.springbootstudy.bean.WxBean.Wxuser;
import lombok.Data;

import javax.persistence.*;
import java.util.ArrayList;
import java.util.List;

@Entity
@Table(name = "ass")
@Data
public class Ass {
    @Id
    @GeneratedValue(strategy= GenerationType.AUTO)
    //社团id
    private int assid;

    private String logo;

    //社团名称
    private String assname;
    //社团简介
    private String assDescription;
    //社团类型
    private String asstype;

    //社团经费
    private double money;
    //优秀社团
    private int isExecllent;
    //社长名称
    private String presidentname;

    //社长学号
    private String presidentId;

    //指导老师姓名
    private String teachername;

    //指导老师联系方式
    private String teacherphone;

    //指导老师所在单位
    private String teacherpost;

    //关联表
    @OneToMany(targetEntity = User.class)
    //从表中的外键+对应的主表中的主键
    @JoinColumn(name="assEntity",referencedColumnName =  "assid")
    private List<User> user;

    public List<User> getUserList(){
        return user;
    }

    public JSONObject getAss(){
        JSONObject jsonObject = new JSONObject();
        jsonObject.put("assId",assid);
        jsonObject.put("logo",logo);
        jsonObject.put("assName",assname);
        jsonObject.put("assType",asstype);
        jsonObject.put("assDescription",assDescription);
        jsonObject.put("money",money);
        jsonObject.put("isExecllent",isExecllent);
        jsonObject.put("presidentName",presidentname);
        jsonObject.put("presidentId",presidentId);
        jsonObject.put("teacherName",teachername);
        jsonObject.put("teacherPhone",teacherphone);
        jsonObject.put("teacherPost",teacherpost);
        return jsonObject;
    }
}


