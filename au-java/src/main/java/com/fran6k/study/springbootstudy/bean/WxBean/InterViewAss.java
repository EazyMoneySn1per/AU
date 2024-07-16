package com.fran6k.study.springbootstudy.bean.WxBean;

import com.fran6k.study.springbootstudy.bean.User;
import lombok.Data;
import org.hibernate.annotations.GenericGenerator;

import javax.persistence.*;
import java.io.Serializable;
import java.util.List;

@Entity
@Data
public class InterViewAss implements Serializable {
    @Id
    @GeneratedValue(generator = "paymentableGenerator")
    @GenericGenerator(name = "paymentableGenerator", strategy = "uuid")
    String id;
    int assId;
    String assName;
    String logo;
    String presidentName;
    String presidentWechat;
    String showMessage;
    String confirmJoinMessage;
    String codeUrl;

    @OneToMany(targetEntity = InterViewUser.class)
    //从表中的外键+对应的主表中的主键
    @JoinColumn(name="submitAssId",referencedColumnName =  "assId")
    private List<InterViewUser> interViewUsers;

}
