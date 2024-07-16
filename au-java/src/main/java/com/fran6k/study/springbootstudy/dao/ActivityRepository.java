package com.fran6k.study.springbootstudy.dao;

import com.fran6k.study.springbootstudy.bean.InReviewBean.Activity;
import com.fran6k.study.springbootstudy.bean.InReviewBean.ActivityComplete;
import com.fran6k.study.springbootstudy.bean.Twitter;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.ArrayList;

public interface ActivityRepository extends JpaRepository<Activity,String> {
    Activity findByAssid(int assid);
    Activity findByAssidAndStepIsNot(int assid,int step);
    ArrayList<Activity> findAllByStepNot(int n);
    @Query(value = "SELECT * FROM `activity` WHERE DATE(create_time) = CURDATE()", nativeQuery = true)
    ArrayList<Activity> findTodayAll();

}
