package com.fran6k.study.springbootstudy.dao;

import com.fran6k.study.springbootstudy.bean.Ass;
import com.fran6k.study.springbootstudy.bean.WxBean.InterViewAss;
import org.springframework.data.jpa.repository.JpaRepository;

public interface InterViewAssRepository  extends JpaRepository<InterViewAss,String> {
    InterViewAss findByAssId(int n);
    InterViewAss findByAssName(String name);
}
