package com.fran6k.study.springbootstudy.utils;

import com.alibaba.fastjson.JSONObject;
import lombok.Data;
import net.sf.json.JSONArray;

@Data
public class ResultModelArrary {
    private int code;
    private String msg;
    private JSONArray data;

    public ResultModelArrary(int code, String msg, JSONArray jsonArray){
        this.code=code;
        this.msg=msg;
        this.data=jsonArray;
    }

    public ResultModelArrary(int code, String msg){
        this.code=code;
        this.msg=msg;
        this.data=null;
    }
}
