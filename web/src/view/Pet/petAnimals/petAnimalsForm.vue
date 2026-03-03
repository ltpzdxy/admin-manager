
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="昵称:" prop="nickname">
    <el-input v-model="formData.nickname" :clearable="true" placeholder="请输入昵称" />
</el-form-item>
        <el-form-item label="品种:" prop="breed">
    <el-input v-model="formData.breed" :clearable="true" placeholder="请输入品种" />
</el-form-item>
        <el-form-item label="是否绝育:" prop="is_neutered">
    <el-switch v-model="formData.is_neutered" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
</el-form-item>
        <el-form-item label="体重:" prop="weight">
    <el-input-number v-model="formData.weight" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="过敏/禁忌:" prop="allergy">
    <el-input v-model="formData.allergy" :clearable="true" placeholder="请输入过敏/禁忌" />
</el-form-item>
        <el-form-item label="照片:" prop="avatar">
    <SelectImage
     v-model="formData.avatar"
     file-type="image"
    />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createPetAnimals,
  updatePetAnimals,
  findPetAnimals
} from '@/api/Pet/petAnimals'

defineOptions({
    name: 'PetAnimalsForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            nickname: '',
            breed: '',
            is_neutered: false,
            weight: 0,
            allergy: '',
            avatar: "",
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPetAnimals({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createPetAnimals(formData.value)
               break
             case 'update':
               res = await updatePetAnimals(formData.value)
               break
             default:
               res = await createPetAnimals(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
