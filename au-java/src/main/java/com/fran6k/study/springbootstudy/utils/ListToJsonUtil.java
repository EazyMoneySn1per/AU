package com.fran6k.study.springbootstudy.utils;

import com.alibaba.fastjson.JSONArray;

import java.util.List;

public class ListToJsonUtil {
    public static<T> JSONArray ListToJson(List<T> s){
        JSONArray jsonArray = new JSONArray();
        for(T string:s){
            jsonArray.add(string);
        }

        return jsonArray;
    }
}
