<template>
  <b-container>
    <br/>
    <div id="app">
      <h1 class="text-center">Droplify - expose your static content</h1>
      <hr/>
      <b-card>
        <b-card-body>
          <b-row>
            <b-col md="12">
              <b-input placeholder="Give your project a name here" v-model="projectName"/>
              <br/>
              <p>
                <span v-if="validName" class="text-success">Valid name</span><span v-else class="text-warning">Invalid name</span>
                Your webapp will live here : <a :href="webAppUrl"><b>{{ webAppUrl }}</b></a>
              </p>
            </b-col>
            <b-col md="12">
              <hr/>
              <form :style="{'display': validName ?  '' : 'none'}" id="my-form" :action="action"
                    class="dropzone">
                <input type="hidden" name="ns" :value="projectName"/>
              </form>
            </b-col>
          </b-row>
        </b-card-body>
      </b-card>
    </div>
  </b-container>
</template>

<script>
import Dropzone from "dropzone";

export default {
  data() {
    return {
      projectName: '',
      files: undefined
    }
  },
  mounted() {
    let myDropzone = new Dropzone(".dropzone", {
      init: function () {
        this.hiddenFileInput.setAttribute("webkitdirectory", true);
        this.hiddenFileInput.setAttribute("name", "lol");
        console.log('owi', this.hiddenFileInput)
      },
      addedfiles: function () {
        console.log('on add file', arguments);
      },
      sending: function () {
        console.log('on sending', arguments)
      }
    });
    myDropzone.on("addedfile", file => {
      console.log(`File added: ${file.name}`);
    });
  },
  computed: {
    webAppUrl() {
      return this.validName ? `${location.protocol}//${this.projectName}.${location.hostname}${location.port ? ':' + location.port : ''}/` : '#'
    },
    validName() {
      return /^[a-z0-9]{1}[a-z0-9-]+[a-z0-9]{1}$/.test(this.projectName)
    },
    action() {
      return process.env.VUE_APP_DEV_MODE === '1' ? 'http://localhost:8020/upload' : '/upload'
    }
  },

}
</script>

<style>
html, body, #app, .container {
  height: 100%;
}

.card {
  height: 90%;
}

.card-body {
  height: 60%;
}

</style>
