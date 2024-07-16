package com.fran6k.study.springbootstudy.dao;

import com.fran6k.study.springbootstudy.bean.InReviewBean.ActivityComplete;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.ArrayList;
import java.util.List;

public interface ActivityCompleteRepository extends JpaRepository<ActivityComplete,String> {
    ActivityComplete findById(int id);
    Page<ActivityComplete> findAll(Pageable Pageable);
    ArrayList<ActivityComplete> findAll();
    Page<ActivityComplete> findAllByAss_AssnameOrAss_AssidOrderByDate(String assName, int assidInt, Pageable Pageable);
}
