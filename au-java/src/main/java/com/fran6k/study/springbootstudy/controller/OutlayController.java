package com.fran6k.study.springbootstudy.controller;

import com.fran6k.study.springbootstudy.bean.Outlay;
import com.fran6k.study.springbootstudy.bean.Twitter;
import com.fran6k.study.springbootstudy.dao.OutlayRepository;
import com.fran6k.study.springbootstudy.utils.ResultModelArrary;
import com.fran6k.study.springbootstudy.utils.UUIDUtils;
import net.sf.json.JSONArray;
import org.apache.tomcat.util.http.fileupload.IOUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.multipart.MultipartFile;

import javax.servlet.http.HttpServletResponse;
import java.io.File;
import java.io.FileInputStream;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Date;

@Controller
public class OutlayController {
    //步骤
    public static int finalstep = 4;

    @Autowired
    OutlayRepository outlayRepository;

    @Value("${upload.outlay.path}")
    private String rootPath;

    //财务报表上传
    @PostMapping("outlay/outlaysubmit")
    @ResponseBody
    public ResultModelArrary twittersubmit(@RequestParam("files") MultipartFile file,
                                           @RequestParam("assid") int assid,
                                           @RequestParam("name") String name,
                                           @RequestParam("step") int step,
                                           @RequestParam("money") double money) throws Exception{

        Outlay outlay = outlayRepository.findByAssidAndStepIsNot(assid,finalstep);

        if(outlay==null){
            outlay = new Outlay();
        }
        //判断是否是二次提交
        if(outlay.getOutlayUrl()!=null){//删除上一次提交的文件
            String prePicUrl = outlay.getOutlayUrl();
            String temPath = rootPath + "/" + prePicUrl;
            File deletefile = new File(temPath);
            deletefile.delete();

        }
        String outlayUrl=this.upLoadOutlay(file);
        outlay.setOutlayUrl(outlayUrl);

        outlay.setAssid(assid);
        outlay.setName(name);
        outlay.setStep(++step);
        outlay.setMoney(money);
        outlay.setBackmsg(null);

        outlayRepository.save(outlay);


        ResultModelArrary resultModelArrary = new ResultModelArrary(20000,"ok");
        return resultModelArrary;
    }


    private String upLoadOutlay(MultipartFile file) throws Exception {
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


    //2.当社长进入页面的时候，获取进度和失败信息
    @GetMapping("outlay/getinfo")
    @ResponseBody
    public ResultModelArrary getinfo(@RequestParam("assid") int assid){
        ResultModelArrary resultModel = new ResultModelArrary(20000,"ok");

        Outlay outlay = outlayRepository.findByAssidAndStepIsNot(assid,finalstep);
        if(outlay!=null){
            resultModel.setData(JSONArray.fromObject(outlay));
        }
        return resultModel;
    }

    //3.审核成功后下一步
    @GetMapping("outlay/nextstep")
    @ResponseBody
    public ResultModelArrary nextstep(@RequestParam("assid") int assid){
        Outlay outlay = outlayRepository.findByAssidAndStepIsNot(assid,finalstep);

        int temstep = outlay.getStep();
        outlay.setStep(++temstep);
        outlayRepository.save(outlay);

        ResultModelArrary resultModelArrary = new ResultModelArrary(20000,"ok");
        return resultModelArrary;
    }
    //4.审核失败的返回信息
    @GetMapping("outlay/setbackmsg")
    @ResponseBody
    public ResultModelArrary setbackmsg(@RequestParam("assid") int assid,
                                        @RequestParam("backmsg") String _backmsg){
        Outlay outlay = outlayRepository.findByAssidAndStepIsNot(assid,finalstep);
//        int temstep = outlay.getStep();
        outlay.setStep(1);
        outlay.setBackmsg(_backmsg);
        outlayRepository.save(outlay);
        ResultModelArrary resultModelArrary = new ResultModelArrary(20000,"ok");
        return resultModelArrary;
    }
    //5.获取等待审核列表
    @GetMapping("outlay/getOutlays")
    @ResponseBody
    public ResultModelArrary getactivities(){
        ArrayList<Outlay> outlays = outlayRepository.findAllByStepNot(finalstep);
        ResultModelArrary resultModelArrary = new ResultModelArrary(20000,"ok",JSONArray.fromObject(outlays));
        return resultModelArrary;
    }

    //6.文件下载
    @GetMapping("outlay/downFile")
    @ResponseBody
    public void downFile(@RequestParam("path") String path, HttpServletResponse response) throws Exception{
        String downPath = rootPath + "/" + path;
        File file = new File(downPath);
        String[] filename = file.getName().split("_");
        FileInputStream fileInputStream = null;
        fileInputStream = new FileInputStream(file);
        response.setContentType("application/force-download");
        response.setCharacterEncoding("UTF-8");
        response.setHeader("Content-Disposition", "attachment;fileName=" +   java.net.URLEncoder.encode(filename[1],"UTF-8"));
        IOUtils.copy(fileInputStream ,response.getOutputStream());
    }
}
