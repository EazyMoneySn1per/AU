package com.fran6k.study.springbootstudy.dao;

import com.fran6k.study.springbootstudy.bean.Outlay;
import com.fran6k.study.springbootstudy.bean.Twitter;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.ArrayList;
import java.util.Date;

public interface OutlayRepository  extends JpaRepository<Outlay,String> {
    Outlay findByAssidAndStepIsNot(int assid, int step);
    ArrayList<Outlay> findAllByStepNot(int n);
    ArrayList<Outlay> findAllByStep(int n);
    Page<Outlay> findAllByStep(int n, Pageable pageable);
    Page<Outlay> findAllByAssidAndStep(int step, int assid, Pageable pageable);

    @Query(value = "SELECT * FROM `outlay` WHERE DATE(create_time) = CURDATE()", nativeQuery = true)
    ArrayList<Outlay> findTodayAll();
}
