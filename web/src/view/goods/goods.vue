<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="商品名称">
          <el-input placeholder="商品名称" v-model="searchInfo.title"></el-input>
        </el-form-item>
        <!-- <el-form-item label="描述">
          <el-input placeholder="描述" v-model="searchInfo.description"></el-input>
        </el-form-item>
        <el-form-item label="api组">
          <el-input placeholder="api组" v-model="searchInfo.apiGroup"></el-input>
        </el-form-item> -->
        <!-- <el-form-item label="请求">
          <el-select clearable placeholder="请选择" v-model="searchInfo.method">
            <el-option
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
              v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item> -->
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="onPullGoods" type="primary">拉取商品</el-button>
        </el-form-item>
        <!-- <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item> -->
      </el-form>
    </div>

    <el-table :data="tableData" @sort-change="sortChange" border stripe @selection-change="handleSelectionChange">
       <el-table-column
        type="selection"
        width="55">
      </el-table-column>
      <el-table-column label="商品编号" prop="id" sortable="custom"></el-table-column>
      <el-table-column label="商品图片" >
        <template slot-scope="scope">
            <CustomPic picType="file" :picSrc="scope.row.mainPic" />
          </template>
      </el-table-column>
      <el-table-column label="商品名称"  prop="title" sortable="custom"></el-table-column>
      <el-table-column label="品牌"  prop="brandName" sortable="custom"></el-table-column>
      <el-table-column label="旗舰店"  prop="shopName" sortable="custom"></el-table-column>
      <el-table-column label="商品价格"  prop="originalPrice" sortable="custom"></el-table-column>
      <el-table-column label="折扣价格"  prop="actualPrice" sortable="custom"></el-table-column>
      <el-table-column label="佣金率"  prop="commissionRate" sortable="custom"></el-table-column>
      <el-table-column label="券总量"  prop="couponTotalNum" sortable="custom"></el-table-column>
      <el-table-column label="开始时间"  prop="couponStartTime" sortable="custom"></el-table-column>
      <el-table-column label="结束时间"  prop="couponEndTime" sortable="custom"></el-table-column>

      <!-- <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editApi(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-button
            @click="deleteApi(scope.row)"
            size="small"
            type="danger"
            icon="el-icon-delete"
          >删除</el-button>
        </template>
      </el-table-column> -->
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <!-- <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="apiForm">
        <el-form-item label="路径" prop="path">
          <el-input autocomplete="off" v-model="form.path"></el-input>
        </el-form-item>
        <el-form-item label="请求" prop="method">
          <el-select placeholder="请选择" v-model="form.method">
            <el-option
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
              v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="api分组" prop="apiGroup">
          <el-input autocomplete="off" v-model="form.apiGroup"></el-input>
        </el-form-item>
        <el-form-item label="api简介" prop="description">
          <el-input autocomplete="off" v-model="form.description"></el-input>
        </el-form-item>
      </el-form>
      <div class="warning">新增Api需要在角色管理内配置权限才可使用</div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog> -->
  </div>
</template>


<script>
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import {
  getGoodsList,
  pullGoods,
} from "@/api/goods";//从哪里引入接口路由
import infoList from "@/mixins/infoList";
import { toSQLLine } from "@/utils/stringFun";
import CustomPic from "@/components/customPic";//引入自定义组件 需注册


export default {
  name: "Goods",
  mixins: [infoList],
  components:{
    CustomPic //注册组件
  },
  data() {
    return {
      deleteVisible:false,
      listApi: getGoodsList,
      dialogFormVisible: false,
      apis:[],
      type: "",
      rules: {
        title: [{ required: true, message: "请输入商品名称", trigger: "blur" }]
        // apiGroup: [
        //   { required: true, message: "请输入组名称", trigger: "blur" }
        // ],
        // method: [
        //   { required: true, message: "请选择请求方式", trigger: "blur" }
        // ],
        // description: [
        //   { required: true, message: "请输入api介绍", trigger: "blur" }
        // ]
      }
    };
  },
  methods: {
    //  选中api
      handleSelectionChange(val) {
        this.apis = val;
        alert(val)
      },
    // 排序
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop);
        this.searchInfo.desc = order == "descending";
      }
      this.getTableData();
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1;
      this.pageSize = 10;
      this.title = this.searchInfo.title
      this.getTableData();
    },
    async onPullGoods(){
      const res = await pullGoods();
      if(res.code==0){
          this.$message({
            type:"success",
            message:res.msg
          })
          this.deleteVisible = false
          this.getTableData()
        }
    },
  },
  filters: {
    // methodFiletr(value) {
    //   const target = methodOptions.filter(item => item.value === value)[0];
    //   // return target && `${target.label}(${target.value})`
    //   return target && `${target.label}`;
    // },
    // tagTypeFiletr(value) {
    //   const target = methodOptions.filter(item => item.value === value)[0];
    //   return target && `${target.type}`;
    // }
  },
  created() {
    this.page = 1;
    this.pageSize = 10;
    this.getTableData();
  }
};
</script>
<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.el-tag--mini {
  margin-left: 5px;
}
.warning {
  color: #dc143c;
}
</style>