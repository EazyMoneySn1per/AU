package com.fran6k.study.springbootstudy.controller;

import com.alibaba.fastjson.JSONObject;
import com.fran6k.study.springbootstudy.bean.Twitter;
import com.fran6k.study.springbootstudy.dao.TwitterRepository;
import com.fran6k.study.springbootstudy.utils.ResultModel;
import com.fran6k.study.springbootstudy.utils.ResultModelArrary;
import com.fran6k.study.springbootstudy.utils.UUIDUtils;
import net.sf.json.JSONArray;
import org.apache.logging.log4j.util.Strings;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.io.FileInputStream;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;

@Controller
public class TwitterController {
    //步骤
    public static int finalstep = 4;

    @Autowired
    TwitterRepository twitterRepository;

    //推文图片上传+请求提交

    @PostMapping("twitter/twittersubmit")
    @ResponseBody
    public ResultModelArrary twittersubmit(@RequestParam("files") MultipartFile[] file,
                                           @RequestParam("assid") int assid,
                                           @RequestParam("name") String name,
                                           @RequestParam("step") int step) throws Exception{
        String picUrl = this.upLoadPicture(file);
        Twitter twitter = twitterRepository.findByAssidAndStepIsNot(assid,finalstep);
        if(twitter==null){
            twitter = new Twitter();
        }
        //判断是否是二次提交
        if(twitter.getPictureUrl()!=null){//删除上一次提交的文件
            String[] prePicUrl = twitter.getPictureUrl().split("`");
            for(int i=1;i<prePicUrl.length;i++){
                String temPath = rootPath + "/" + prePicUrl[i];
                File deletefile = new File(temPath);
                deletefile.delete();
            }
        }
        twitter.setPictureUrl(picUrl);
        twitter.setAssid(assid);
        twitter.setName(name);
        twitter.setStep(++step);
        twitter.setBackmsg(null);

        twitterRepository.save(twitter);


        ResultModelArrary resultModelArrary = new ResultModelArrary(20000,"ok");
        return resultModelArrary;
    }

    //图片文件根目录地址
//    private  final static String rootPath = "/Users/frankkkkkk/Desktop/twitter";
    @Value("${upload.picture.path}")
    private String rootPath;

    //将图片存入本地服务器
    private String upLoadPicture(MultipartFile[] file) throws Exception {

        String splicePath="";
        for(MultipartFile files:file){
            //使用uuid工具
            String uuid = UUIDUtils.getUUID()+"";
            //创建日期
            SimpleDateFormat simpleDateFormat = new SimpleDateFormat("yyyy-MM-dd");
            Date date = new Date();
            String dateStr = simpleDateFormat.format(date);

            //创建时间目录
            String storagePathFather=rootPath+"/"+dateStr;

            File fileDir = new File(storagePathFather);
            if (!fileDir.exists() && !fileDir.isDirectory())
                fileDir.mkdirs();

            //数据库中存储的地址
            String dbStoragePath  = dateStr + "/"+uuid+"_"+files.getOriginalFilename();
            //实际物理存储地址
            String storagePath = rootPath +"/"+dbStoragePath;
            files.transferTo(new File(storagePath));

            //将所有文件的地址拼接成一个地址
            if(splicePath!=""){
                splicePath=splicePath + "`" + dbStoragePath;
            }
            else
                splicePath=dbStoragePath;

        }
        return splicePath;
    }

    //推文图片
    @GetMapping("twitter/downFile")
    @ResponseBody
    public byte[] downFile(@RequestParam("path") String path) throws Exception{
        if (path.contains("../") || path.contains("./")) {
            return null;
        }
        if (path != "") {
            File file = new File(rootPath + "/" + path);
            FileInputStream inputStream = new FileInputStream(file);
            byte[] bytes = new byte[inputStream.available()];
            inputStream.read(bytes,0,inputStream.available());
            return bytes;
        }
        return null;
    }

    //2.当社长进入页面的时候，获取推文进度和失败信息
    @GetMapping("twitter/getinfo")
    @ResponseBody
    public ResultModelArrary getinfo(@RequestParam("assid") int assid){
        ResultModelArrary resultModel = new ResultModelArrary(20000,"ok");

        Twitter twitter = twitterRepository.findByAssidAndStepIsNot(assid,finalstep);
        if(twitter!=null){
//            JSONObject jsonObject = new JSONObject();
//            jsonObject.put("name",twitter.getName());
//            jsonObject.put("pictureUrl",twitter.getPictureUrl());
//            jsonObject.put("step",twitter.getStep());
//            jsonObject.put("backmsg",twitter.getBackmsg());
//            jsonObject.put("")
            resultModel.setData(JSONArray.fromObject(twitter));
        }
        return resultModel;
    }

    //3.审核成功后下一步
    @GetMapping("twitter/nextstep")
    @ResponseBody
    public ResultModelArrary nextstep(@RequestParam("assid") int assid){
        Twitter twitter = twitterRepository.findByAssidAndStepIsNot(assid,finalstep);

        int temstep = twitter.getStep();
        twitter.setStep(++temstep);
        twitterRepository.save(twitter);

        ResultModelArrary resultModelArrary = new ResultModelArrary(20000,"ok");
        return resultModelArrary;
    }
    //4.审核失败的返回信息
    @GetMapping("twitter/setbackmsg")
    @ResponseBody
    public ResultModelArrary setbackmsg(@RequestParam("assid") int assid,
                                        @RequestParam("backmsg") String _backmsg){
        Twitter twitter = twitterRepository.findByAssidAndStepIsNot(assid,finalstep);
//        int temstep = twitter.getStep();
        twitter.setStep(1);
        twitter.setBackmsg(_backmsg);
        twitterRepository.save(twitter);
        ResultModelArrary resultModelArrary = new ResultModelArrary(20000,"ok");
        return resultModelArrary;
    }
    //5.获取等待审核列表
    @GetMapping("twitter/gettwitters")
    @ResponseBody
    public ResultModelArrary getactivities(){
        ArrayList<Twitter> twitters = twitterRepository.findAllByStepNot(finalstep);
        ResultModelArrary resultModelArrary = new ResultModelArrary(20000,"ok",JSONArray.fromObject(twitters));
        return resultModelArrary;
    }
}
