package com.fran6k.study.springbootstudy.bean.InReviewBean;

import lombok.Data;
import org.hibernate.annotations.GenericGenerator;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import javax.persistence.*;
import java.util.Date;

@Entity
@Table(name = "activity")
@Data
@EntityListeners(AuditingEntityListener.class)

public class Activity {
    @Id
    @GeneratedValue(generator = "paymentableGenerator")
    @GenericGenerator(name = "paymentableGenerator", strategy = "uuid")
    private String id;

    @Column(name = "name")
    private String name;


    @Column(name = "assid")
    private int assid;

    @Column(name = "date")
    private Date date;

    @Column(name = "step")
    private int step;

    @Column(name = "file1")
    private String file1;

    @Column(name = "file2")
    private String file2;

    @Column(name = "backmsg")
    private String backmsg;

    @Column(name = "description")
    private String description;

    @CreatedDate//自动添加创建时间的注解
    private Date createTime;
    @LastModifiedDate//自动添加更新时间的注解
    private Date updateTime;


}
