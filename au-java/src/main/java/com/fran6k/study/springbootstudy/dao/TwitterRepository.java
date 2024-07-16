package com.fran6k.study.springbootstudy.dao;

import com.fran6k.study.springbootstudy.bean.Outlay;
import com.fran6k.study.springbootstudy.bean.Twitter;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.ArrayList;
import java.util.Date;

public interface TwitterRepository extends JpaRepository<Twitter,String> {
    Twitter findByAssidAndStepIsNot(int assid,int step);
    ArrayList<Twitter> findAllByStepNot(int n);
    ArrayList<Twitter> findAllByStep(int n);

    @Query(value = "SELECT * FROM `twitter` WHERE DATE(create_time) = CURDATE()", nativeQuery = true)
    ArrayList<Twitter> findTodayAll();

    Page<Twitter> findAllByStep(int n, Pageable pageable);
    Page<Twitter> findAllByAssidAndStep(int step, int assid, Pageable pageable);

}
