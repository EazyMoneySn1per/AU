package com.fran6k.study.springbootstudy.bean;

import lombok.Data;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "outlay")
@Data
@EntityListeners(AuditingEntityListener.class)
public class Outlay {
    @Id
    @GeneratedValue(strategy= GenerationType.AUTO)
    private int id;

    //报销活动
    private String name;
    //本次报销金额
    private double money;

    //所属社团id
    private int assid;

    //上传的url
    private String outlayUrl;

    private int step;
    private String backmsg;

    @CreatedDate//自动添加创建时间的注解
    private Date createTime;
    @LastModifiedDate//自动添加更新时间的注解
    private Date updateTime;

}
