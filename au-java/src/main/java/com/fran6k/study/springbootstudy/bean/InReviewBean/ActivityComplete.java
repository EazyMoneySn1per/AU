package com.fran6k.study.springbootstudy.bean.InReviewBean;

import com.alibaba.fastjson.JSONObject;
import com.fran6k.study.springbootstudy.bean.Ass;
import lombok.Data;
import net.sf.json.JSONArray;
import org.springframework.data.annotation.CreatedDate;
import org.springframework.data.annotation.LastModifiedDate;
import org.springframework.data.jpa.domain.support.AuditingEntityListener;

import javax.persistence.*;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.List;

@Entity
@Table(name = "activity_complete")
@Data
@EntityListeners(AuditingEntityListener.class)
public class ActivityComplete {
    @Id
    @GeneratedValue(strategy= GenerationType.AUTO)
    private int id;
    private String activityName;
    private String twitter;
    private Date date;
    private String file1;
    private String file2;
    private String description;
    @CreatedDate//自动添加创建时间的注解
    private Date createTime;
    @LastModifiedDate//自动添加更新时间的注解
    private Date updateTime;

    @ManyToOne(targetEntity= Ass.class)
    @JoinColumn(name="ass",referencedColumnName = "assid")
    private Ass ass;

    public JSONObject getInfo(){
        String strDateFormat = "yyyy-MM-dd HH:mm:ss";
        SimpleDateFormat sdf = new SimpleDateFormat(strDateFormat);

        JSONObject jsonObject = new JSONObject();
        jsonObject.put("id",id);
        jsonObject.put("activityName",activityName);
        jsonObject.put("date",sdf.format(date));
        jsonObject.put("twitter",twitter);
        jsonObject.put("file1",file1);
        jsonObject.put("file2",file2);
        jsonObject.put("description",description);
        jsonObject.put("assName",ass.getAssname());
        jsonObject.put("assId",ass.getAssid());
        jsonObject.put("createTime",sdf.format(createTime));
        jsonObject.put("updateTime",sdf.format(updateTime));
        return jsonObject;
    }
}
