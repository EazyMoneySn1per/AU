package com.fran6k.study.springbootstudy.controller;

import com.fran6k.study.springbootstudy.bean.Ass;
import com.fran6k.study.springbootstudy.bean.WxBean.InterViewAss;
import com.fran6k.study.springbootstudy.dao.AssRepository;
import com.fran6k.study.springbootstudy.dao.InterViewAssRepository;
import com.fran6k.study.springbootstudy.utils.ResultModel;
import com.fran6k.study.springbootstudy.utils.ResultModelArrary;
import com.fran6k.study.springbootstudy.utils.UUIDUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.Map;

@Controller
public class AMinisterController {

    @Autowired
    AssRepository assRepository;
    @Autowired
    InterViewAssRepository interViewAssRepository;

    @Value("${upload.assLogo.path}")
    private String assLogoPath;

    @PostMapping("AMinister/uploadAssLogo")
    @ResponseBody
    public ResultModel uploadAssLogo(@RequestParam("files") MultipartFile file,
                                     @RequestParam("assId") int assId) throws Exception{

        Ass ass = assRepository.findByAssid(assId);
        InterViewAss interViewAss = interViewAssRepository.findByAssId(assId);
        //使用uuid工具
        String uuid = UUIDUtils.getUUID()+"";
        //创建日期
        SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd");
        Date date = new Date();
        String dateStr = simpleDateFormat.format(date);
        String newName = uuid + "_" + file.getOriginalFilename();
        String storagePath = assLogoPath + "/" + newName;
        ass.setLogo(newName);
        interViewAss.setLogo(newName);
        file.transferTo(new File(storagePath));
        assRepository.save(ass);
        interViewAssRepository.save(interViewAss);
        return new ResultModel(20000,"");
    }

    @PostMapping("AMinister/setAssDescription")
    @ResponseBody
    public ResultModel setAssDescription(@RequestBody Map map) throws Exception{
        int assId = (int) map.get("assId");
        String assDescription = map.get("assDescription").toString();
        Ass ass = assRepository.findByAssid(assId);
        ass.setAssDescription(assDescription);
        assRepository.save(ass);
        return new ResultModel(20000,"");
    }
}
