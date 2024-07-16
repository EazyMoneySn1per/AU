package com.fran6k.study.springbootstudy.dao;

import com.fran6k.study.springbootstudy.bean.WxBean.InterViewAss;
import com.fran6k.study.springbootstudy.bean.WxBean.InterViewUser;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.ArrayList;
import java.util.Optional;

public interface InterViewUserRepository extends JpaRepository<InterViewUser,String> {
    ArrayList<InterViewUser> findAllByStudentId(String studentId);
    InterViewUser findByStudentIdAndInterViewAss(String studentId,InterViewAss interViewAss);

//    @Query("from InterViewUser u  where u.interViewStatus = :status ")
//    ArrayList<InterViewUser> findAllByInterViewStatus(@Param("status") InterViewEnum interViewEnum);

//    @Query(value = "select * from inter_view_user where 1=1 " +
//            "and (?1 is null or inter_view_status = ?1)",nativeQuery = true)
//    Page<InterViewUser> searchAllBy(InterViewEnum interViewEnum, PageRequest pageRequest);
}
