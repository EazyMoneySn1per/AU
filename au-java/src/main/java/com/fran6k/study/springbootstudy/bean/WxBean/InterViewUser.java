package com.fran6k.study.springbootstudy.bean.WxBean;

import com.alibaba.fastjson.JSONObject;
import com.fasterxml.jackson.annotation.JsonFormat;
import com.fran6k.study.springbootstudy.bean.Ass;
import lombok.Data;
import org.hibernate.annotations.GenericGenerator;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;

import javax.persistence.*;
import java.text.SimpleDateFormat;
import java.util.Date;

@Entity
@Data
public class InterViewUser {
    @Id
    @GeneratedValue(generator = "paymentableGenerator")
    @GenericGenerator(name = "paymentableGenerator", strategy = "uuid")
    String id;
    String name;
    String studentId;
    String description;
    String phoneNum;
    String wxNum;
    //backMessage被废除
    String backMessage;
    String sex;
    //按钮控制，1为社团操作，2为学生操作，0为停用
    int buttonControl;
    @Enumerated(EnumType.ORDINAL)
    InterViewEnum interViewStatus = InterViewEnum.Stage_One_WaitB;
    @ManyToOne(targetEntity= InterViewAss.class)
    @JoinColumn(name="submitAssId",referencedColumnName = "assId")
    InterViewAss interViewAss;
    @CreatedDate//自动添加创建时间的注解
    @JsonFormat(timezone = "GMT+8", pattern = "yyyy-MM-dd HH:mm:ss")
    private Date createTime;
    @LastModifiedDate//自动添加更新时间的注解
    @JsonFormat(timezone = "GMT+8", pattern = "yyyy-MM-dd HH:mm:ss")
    private Date updateTime;

    public InterViewUser() {}
    public InterViewUser(String studentId,
                         String name,
                         String sex,
                         String description,
                         InterViewAss interViewAss,
                         String phoneNum,
                         String wxNum){
        this.backMessage="";
        this.name=name;
        this.sex=sex;
        this.studentId=studentId;
        this.description=description;
        this.interViewAss=interViewAss;
        this.wxNum = wxNum;
        this.phoneNum = phoneNum;
        this.buttonControl=1;
    }
    public JSONObject getInfo(){
        JSONObject jsonObject = new JSONObject();
        jsonObject.put("id",id);
        jsonObject.put("studentId",studentId);
        jsonObject.put("description",description);
        jsonObject.put("sex",sex);
        jsonObject.put("studentName",name);
        jsonObject.put("wxNum",wxNum);
        jsonObject.put("phoneNum",phoneNum);
        jsonObject.put("assName",interViewAss.getAssName());
        jsonObject.put("assLogo",interViewAss.getLogo());
        jsonObject.put("presidentName",interViewAss.getPresidentName());
        jsonObject.put("presidentWechat",interViewAss.getPresidentWechat());
//        backMessage已经被废除，返回的backMessage为连表查询社团对应的message
//        jsonObject.put("backMessage",backMessage);
        jsonObject.put("interViewStatus",interViewStatus.getStep());
        jsonObject.put("interViewStatusMessage",interViewStatus.getMessage());
        jsonObject.put("buttonControl",buttonControl);
//        SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
//        jsonObject.put("createTime",simpleDateFormat.format(createTime));
//        jsonObject.put("updateTime",simpleDateFormat.format(updateTime));
        return jsonObject;
    }
}
