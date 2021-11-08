# communience-store

## 初回実行

1. 入ってない人はNode.jsを入れる
1. 自分のPCでPowerShell等のbashを開いて`mmunience-store`まで移動  
1. `npm install`  
    を実行
1. `docker-compose up --build`  
   を実行 
1. ```bash
    commuen_front | ✔ Builder initialized
    commuen_front | ✔ Nuxt files generated
    commuen_front | ℹ Compiling Client
    commuen_front | ✔ Client: Compiled successfully in 15.16s
    commuen_front | ℹ Waiting for file changes
    commuen_front | ℹ Memory usage: 311 MB (RSS: 416 MB)
    commuen_front | ℹ Listening on: http://172.19.0.2:8080/
    ```
1. ↑のように表示されたらブラウザを開いて`http://localhost:8080`にアクセス

## 2回目以降
1. 自分のPCでPowerShell等のbashを開いて`communience-store`まで移動
2. `docker-compose up --build`  
   を実行 
1. ```bash
    commuen_front | ✔ Builder initialized
    commuen_front | ✔ Nuxt files generated
    commuen_front | ℹ Compiling Client
    commuen_front | ✔ Client: Compiled successfully in 15.16s
    commuen_front | ℹ Waiting for file changes
    commuen_front | ℹ Memory usage: 311 MB (RSS: 416 MB)
    commuen_front | ℹ Listening on: http://172.19.0.2:8080/
    ```
1. ↑のように表示されたらブラウザを開いて`http://localhost:8080`にアクセス

---
    



## Build Setup

```bash
# install dependencies
$ npm install

# serve with hot reload at localhost:3000
$ npm run dev

# build for production and launch server
$ npm run build
$ npm run start

# generate static project
$ npm run generate
```

For detailed explanation on how things work, check out the [documentation](https://nuxtjs.org).

## Special Directories

You can create the following extra directories, some of which have special behaviors. Only `pages` is required; you can delete them if you don't want to use their functionality.

### `assets`

The assets directory contains your uncompiled assets such as Stylus or Sass files, images, or fonts.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/assets).

### `components`

The components directory contains your Vue.js components. Components make up the different parts of your page and can be reused and imported into your pages, layouts and even other components.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/components).

### `layouts`

Layouts are a great help when you want to change the look and feel of your Nuxt app, whether you want to include a sidebar or have distinct layouts for mobile and desktop.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/layouts).


### `pages`

This directory contains your application views and routes. Nuxt will read all the `*.vue` files inside this directory and setup Vue Router automatically.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/get-started/routing).

### `plugins`

The plugins directory contains JavaScript plugins that you want to run before instantiating the root Vue.js Application. This is the place to add Vue plugins and to inject functions or constants. Every time you need to use `Vue.use()`, you should create a file in `plugins/` and add its path to plugins in `nuxt.config.js`.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/plugins).

### `static`

This directory contains your static files. Each file inside this directory is mapped to `/`.

Example: `/static/robots.txt` is mapped as `/robots.txt`.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/static).

### `store`

This directory contains your Vuex store files. Creating a file in this directory automatically activates Vuex.

More information about the usage of this directory in [the documentation](https://nuxtjs.org/docs/2.x/directory-structure/store).
