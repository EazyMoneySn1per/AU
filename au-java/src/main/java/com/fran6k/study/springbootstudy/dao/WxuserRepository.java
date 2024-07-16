package com.fran6k.study.springbootstudy.dao;

import com.fran6k.study.springbootstudy.bean.Ass;
import com.fran6k.study.springbootstudy.bean.WxBean.Wxuser;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;

import java.util.ArrayList;
import java.util.List;

public interface WxuserRepository extends JpaRepository<Wxuser,String> {
    Wxuser findByOpenId(String openid);
    Wxuser findByStudentId(String id);
    List<Wxuser> findAllByFirstAssOrSecondAss(Ass assOne, Ass assSecond);


    @Query(value = "select * from wxuser where 1=1 " +
            "and (?1 is null or nickname = ?1) " +
            "and (?2 is null or real_name = ?2) " +
            "and (?3 is null or student_id = ?3) " +
            "and (?4 is null or we_chat_id = ?4) " +
            "and (?5 is null or phone_num = ?5) " +
            "and (?6 is null or ass_entity_first = ?6 or ass_entity_second = ?6)", nativeQuery = true)
    Page<Wxuser> searchAllBy(String nickname, String realName, String studentId, String weChatId, String phoneNum, String assid, PageRequest pageRequest);
}
