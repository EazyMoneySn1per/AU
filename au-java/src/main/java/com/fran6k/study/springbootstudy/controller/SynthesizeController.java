package com.fran6k.study.springbootstudy.controller;

import com.fran6k.study.springbootstudy.bean.Outlay;
import com.fran6k.study.springbootstudy.bean.SynthesizeSubmit;
import com.fran6k.study.springbootstudy.dao.AssRepository;
import com.fran6k.study.springbootstudy.dao.SynthesizSubmitRepository;
import com.fran6k.study.springbootstudy.utils.ResultModelArrary;
import com.fran6k.study.springbootstudy.utils.UUIDUtils;
import net.sf.json.JSONArray;
import net.sf.json.JSONObject;
import net.sf.json.JsonConfig;
import net.sf.json.util.CycleDetectionStrategy;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.io.IOException;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;

@Controller
public class SynthesizeController {
    @Autowired
    SynthesizSubmitRepository synthesizSubmitRepository;
    @Autowired
    AssRepository assRepository;


    @Value("${upload.synthesizesubmit.path}")
    private String rootPath;

    @PostMapping("synthesize/synthesizesubmit")
    @ResponseBody
    public ResultModelArrary synthesizesubmit(@RequestParam("files") MultipartFile file,
                                              @RequestParam("assid") int assid,
                                              @RequestParam("description") String description,
                                              @RequestParam("type") int type,
                                              @RequestParam("name") String name) throws Exception{

        SynthesizeSubmit synthesizeSubmit = new SynthesizeSubmit();


//        if(synthesizeSubmit==null){
//            synthesizeSubmit = new SynthesizeSubmit();
//        }
        //判断是否是二次提交
//        if(synthesizeSubmit.getFileUrl()!=null){//删除上一次提交的文件
//            String prePicUrl = outlay.getOutlayUrl();
//            String temPath = rootPath + "/" + prePicUrl;
//            File deletefile = new File(temPath);
//            deletefile.delete();
//
//        }
        synthesizeSubmit.setAss(assRepository.findByAssid(assid));
        synthesizeSubmit.setDescription(description);
        synthesizeSubmit.setName(name);
        synthesizeSubmit.setType(type);
        String fileUrl=this.upLoadFile(file);
        synthesizeSubmit.setFileUrl(fileUrl);

        synthesizSubmitRepository.save(synthesizeSubmit);
//        synthesizeSubmit.setId(UUIDUtils.getUUID());


//        String outlayUrl=this.upLoadOutlay(file);
//        outlay.setOutlayUrl(outlayUrl);
//
//        outlay.setAssid(assid);
//        outlay.setName(name);
//        outlay.setStep(++step);
//        outlay.setMoney(money);
//        outlay.setBackmsg(null);
//
//        outlayRepository.save(outlay);
//

        ResultModelArrary resultModelArrary = new ResultModelArrary(20000,"ok");
        return resultModelArrary;
    }

    private String upLoadFile(MultipartFile file) throws IOException {
        String uuid = UUIDUtils.getUUID()+"";
        SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd");
        Date date = new Date();
        String dateStr = simpleDateFormat.format(date);

        //创建时间目录
        String storagePathFather=rootPath+"/"+dateStr;
        File fileDir = new File(storagePathFather);
        if (!fileDir.exists() && !fileDir.isDirectory())
            fileDir.mkdirs();

        //数据库存储地址
        String dbStoragePath  = dateStr + "/"+uuid+"_"+file.getOriginalFilename();

        //实际物理存储地址
        String storagePath = rootPath + "/" + dbStoragePath;
        file.transferTo(new File(storagePath));

        return dbStoragePath;
    }

    @GetMapping("synthesize/synthesizegetlist")
    @ResponseBody
    public ResultModelArrary getlist(@RequestParam("assid") int assid){
        ArrayList<SynthesizeSubmit> synthesizeSubmits = synthesizSubmitRepository.findByAss_AssidOrderByCreateTimeDesc(assid);
        //防止死循环
        JsonConfig jsonConfig = new JsonConfig();
        jsonConfig.setCycleDetectionStrategy(CycleDetectionStrategy.LENIENT);
        jsonConfig.setExcludes(new String[]{"ass"});  //此处是亮点，只要将所需忽略字段加到数组中即可
        jsonConfig.setIgnoreDefaultExcludes(false);  //设置默认忽略
        JSONArray jsonArray =JSONArray.fromObject(synthesizeSubmits, jsonConfig);

        ResultModelArrary resultModelArrary = new ResultModelArrary(20000,"ok",jsonArray);
        return resultModelArrary;
    }
}
