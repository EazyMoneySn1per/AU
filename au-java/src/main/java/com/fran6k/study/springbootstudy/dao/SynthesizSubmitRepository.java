package com.fran6k.study.springbootstudy.dao;

import com.fran6k.study.springbootstudy.bean.Ass;
import com.fran6k.study.springbootstudy.bean.SynthesizeSubmit;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.ArrayList;
import java.util.List;

public interface SynthesizSubmitRepository extends JpaRepository<SynthesizeSubmit,String> {
//    SynthesizeSubmit findByAssid(int n);
//    SynthesizeSubmit findByAss_Assid(int n);
    ArrayList<SynthesizeSubmit> findByAss_AssidOrderByCreateTimeDesc(int n);
}
