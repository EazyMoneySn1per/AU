package com.fran6k.study.springbootstudy.bean.Vo;

import lombok.Data;

import java.util.List;

@Data
public class VoUserInfo {
    private List<String> roles;
    private String introduction;
    private String avatar;
    private String name;
    private int assid;
}
