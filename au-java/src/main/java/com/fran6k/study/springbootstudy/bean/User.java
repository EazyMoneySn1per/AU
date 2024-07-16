package com.fran6k.study.springbootstudy.bean;

import com.alibaba.fastjson.JSONObject;
import lombok.Data;
import org.hibernate.annotations.GenericGenerator;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;


import javax.persistence.*;

@Entity
@Table(name = "user")
@Data
public class User {
    @Id
    @GeneratedValue(generator = "paymentableGenerator")
    @GenericGenerator(name = "paymentableGenerator", strategy = "uuid")
    private String id;

    @Column(name = "name")
    private String name;

    @Column(name = "password")
    private String password;

    @Column(name = "role")
    private String role;

    @Column(name = "assid")
    private int assid;

    //指定关联的实体，一个人可以加多个社团，一对多
    @ManyToOne(targetEntity=Ass.class)
    @JoinColumn(name="assEntity",referencedColumnName = "assid")
    private Ass ass;

    public Ass getass(){
        return ass;
    }
    public String getName(){
        return name;
    }

    public JSONObject getInfo(){
        JSONObject jsonObject = new JSONObject();
        jsonObject.put("id",this.getId());
        jsonObject.put("name",this.getName());
        jsonObject.put("password",this.getPassword());
        jsonObject.put("role",this.getRole());
        jsonObject.put("assId",this.getAssid());
        if(this.getass()!=null)
            jsonObject.put("assName",this.getass().getAssname());
        else
            jsonObject.put("assName","无");
        return jsonObject;
    }
}
