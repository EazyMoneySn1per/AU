package com.fran6k.study.springbootstudy.bean;

import com.fasterxml.jackson.annotation.JsonFormat;
import lombok.Data;
import org.hibernate.annotations.GenericGenerator;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "synthesizesubmit")
@Data
@EntityListeners(AuditingEntityListener.class)
public class SynthesizeSubmit {
    @Id
    @GeneratedValue(generator = "paymentableGenerator")
    @GenericGenerator(name = "paymentableGenerator", strategy = "uuid")
    private String id;

    //描述
    private String description;

    //类型
    //1.年审资料提交，2.内部信息更换，3外聘老师
    private int type;

    //提交人
    private String name;

    //文件地址
    private String fileUrl;

    //返回信息
    private String backmsg;

    //受理状况
    private int isHandle;

    @ManyToOne(targetEntity=Ass.class)
    @JoinColumn(name="assEntity",referencedColumnName = "assid")
    private Ass ass;

    @CreatedDate//自动添加创建时间的注解
    @JsonFormat(timezone = "GMT+8", pattern = "yyyy-MM-dd HH:mm:ss")
    private Date createTime;

    @LastModifiedDate//自动添加更新时间的注解
    @JsonFormat(timezone = "GMT+8", pattern = "yyyy-MM-dd HH:mm:ss")
    private Date updateTime;

    public SynthesizeSubmit(){
        this.isHandle=0;
    }
}
