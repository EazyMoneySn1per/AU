package com.fran6k.study.springbootstudy.dao;

import com.fran6k.study.springbootstudy.bean.Ass;
import org.springframework.data.domain.Sort;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface AssRepository  extends JpaRepository<Ass,String> {
//    List<Ass> findByUser();
    Ass findByAssid(int n);
    Ass findByAssname(String name);
    List<Ass> findAll();
    List<Ass> findAllByAssidIsNot(int i);
    List<Ass> findAllByIsExecllent(int n);

}
