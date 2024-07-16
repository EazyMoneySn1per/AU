package com.fran6k.study.springbootstudy.dao;

import com.fran6k.study.springbootstudy.bean.User;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRepository extends JpaRepository<User,String> {
    User findByName(String name);
    User findByNameAndPassword(String name,String password);
    Page<User> findAllByRole(String role, Pageable pageable);
}
