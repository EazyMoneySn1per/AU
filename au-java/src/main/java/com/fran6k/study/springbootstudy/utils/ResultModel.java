package com.fran6k.study.springbootstudy.utils;

import com.alibaba.fastjson.JSONObject;
import lombok.Data;
import net.sf.json.JSONArray;

@Data
public class ResultModel {
    /**
     * code状态码
     */
    private int code;
    /**
     * 返回信息
     */
    private String msg;
    /**
     * 返回数据
     */
    private JSONObject data;

    public ResultModel(int code, String msg, JSONObject jsonObject){
        this.code=code;
        this.msg=msg;
        this.data=jsonObject;
    }

    public ResultModel(int code, String msg){
        this.code=code;
        this.msg=msg;
        this.data=null;
    }
}
