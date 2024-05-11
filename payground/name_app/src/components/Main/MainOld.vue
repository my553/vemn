<template>
  <div class="Titles">
    <div class="mainMenu">
      <div class="main">
        <b class="mainTitle">Главная</b>
      </div>
      <div class="statistic">
        <b class="statisticTitle">Статистика</b>
      </div>
      <div class="reports">
        <b class="reportsTitle">Отчеты</b>
      </div>
      <div class="analytics">
        <b class="analyticsTitle">Аналитика</b>
      </div>
    </div>
  </div>
  <div class="Sidebar">
    <div class="logo">
      <img class="groupIcon" alt="" src="../../images/logo-main.svg" />
      <div class="title">VULNERABILITY EXPLORER & MANAGEMENT</div>
    </div>
    <div class="nav">
      <button class="home">
        <img class="homeImg" src="../../images/main.svg" alt="" />
        <div class="text">Главная</div>
      </button>
      <button class="stat">
        <img class="statImg" src="../../images/statistic.svg" alt="" />
        <div class="text">Статистика</div>
      </button>
      <button class="reports">
        <img class="repImg" src="../../images/reports.svg" alt="" />
        <div class="text">Отчеты</div>
      </button>
      <button class="analytics">
        <img class="analImg" src="../../images/analytics.svg" alt="" />
        <div class="text">Аналитика</div>
      </button>
    </div>
    <div class="account">
      <button class="accImg">
        <img src="../../images/nav-profile.svg" alt="" />
      </button>
      <div class="name">{{ this.$data.accFIO }}</div>
      <button class="leaveImg">
        <img src="../../images/nav-exit.svg" alt="" />
      </button>
    </div>
  </div>
  <div class="MainInfo">
    <div class="allStats">
      <div class="allNoRes">
        <b class="title"
          >Общее количество неактивных ресурсов -
          <span class="count">{{ this.deactivateResource }}</span>
          <span class="time">Данные актуальны на {{ this.date }} 3:00</span>
        </b>
      </div>
      <div class="allRes">
        <b class="title"
          >Общее количество ресурсов -
          <span class="count">{{ this.resources }}</span>
          <span class="time">Данные актуальны на {{ this.date }} 3:00</span>
        </b>
      </div>
      <div class="Owners">
        <b class="title"
          >Клиентов -
          <span class="count">{{ this.owners }}</span>
          <span class="time">Данные актуальны на {{ this.date }} 3:00</span>
        </b>
      </div>
      <div class="WAF">
        <b class="title">
          Ресурсов за WAF -
          <span class="count">{{ this.waf }}</span>
          <span class="time">Данные актуальны на {{ this.date }} 3:00</span>
        </b>
      </div>
    </div>
    <div class="weekStats">
      <div class="title">
        <b class="text">Результаты недели</b>
      </div>
      <div class="results">
        <div class="past">
          <span class="titlePast">Прошлая неделя</span>
          <span class="newNoRes"
            >Новые неактивные ресурсы ({{ this.lastNoResCount }})</span
          >
          <span class="newWAF"
            >Новые ресурсы за WAF ({{ this.lastNewWafCount }})</span
          >
          <div class="line1"></div>
          <div class="line2"></div>
          <div class="line3"></div>
        </div>
        <div class="current">
          <span class="titlePast">Текущая неделя</span>
          <span class="newNoRes"
            >Новые неактивные ресурсы ({{ this.currNoResCount }})</span
          >
          <span class="newWAF"
            >Новые ресурсы за WAF ({{ this.currNewWafCount }})</span
          >
          <div class="line1"></div>
          <div class="line2"></div>
          <div class="line3"></div>
        </div>
      </div>
    </div>
    <div class="certs">
      <div class="title">
        <b class="text"
          >Сертификаты, срок действия которых закончится в ближайшие 2 месяца
          <span class="count">({{ this.num }})</span></b
        >
      </div>
      <div class="results">
        <div class="head">
          <b class="resources">Ресурс</b><b class="date">Дата</b>
        </div>
        <div class="data__list">
          <ul>
            <li v-for="item in current">
              <div class="data__content">
                <p class="data__content-resource">{{ item.resource }}</p>
                <p class="data__content-date">{{ item.date }}</p>
              </div>
            </li>
            <li v-for="item in next">
              <div class="data__content">
                <p class="data__content-resource">{{ item.resource }}</p>
                <p class="data__content-date">{{ item.date }}</p>
              </div>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
  <div class="Statistic">
    <div class="checkResource">
      <b class="title">Получить информацию о ресурсе</b>
      <div class="inputBlock">
        <input type="text" class="inputField" placeholder="Ресурс" />
        <div class="button"><span class="text">Проверить ресурс</span></div>
      </div>
    </div>
    <div class="filter">
      <div class="addFilter">
        <img class="filterImg" src="../../images/add-filter.svg" alt="" />
        <span class="text">Добавить фильтр</span>
      </div>
      <div class="delFilter">
        <img class="cross" src="../../images/close.svg" alt="" />
        <div class="text">Сбросить фильтр</div>
      </div>
    </div>
    <div class="statisticContainer">
      <div class="head">
        <span class="count">Найдено: 450</span>
        <span class="rowsCount">10 строк</span>
      </div>
      <div class="table">
        <div class="tableHead">
          <span class="resource">Ресурс</span>
          <span class="ip">IP</span>
          <span class="resolve">Статус активности</span>
          <span class="waf">WAF</span>
          <span class="ssl">SSL</span>
          <span class="date">Дата</span>
        </div>
        <div id="rowsContainer"></div>
      </div>
      <div class="pageRoll">
        <button class="rollBack">
          <img src="../../images/resource-open-input.svg" alt="" />
        </button>
        <div class="pages">
          <button class="page1">1</button>
          <button class="page2">2</button>
          <button class="page3">3</button>
          <button class="page4">4</button>
          <span class="middle">...</span>
          <button class="page10">10</button>
        </div>
        <button class="rollForward">
          <img src="../../images/resource-open-input.svg" alt="" />
        </button>
      </div>
    </div>
  </div>
  <div class="Reports">
    <div class="current">
      <div class="title">
        <span class="text">Текущая неделя</span>
      </div>
      <div class="noRes">
        <b class="head">Неактивные ресурсы</b>
        <div class="table">
          <div class="head">
            <span class="resource">Ресурс</span>
            <span class="date">Дата</span>
          </div>
          <div class="data">
            <div class="row1">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
            <div class="row2">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
            <div class="row3">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
            <div class="row4">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
          </div>
        </div>
      </div>
      <div class="newWAF">
        <b class="head">Новые ресурсы за WAF</b>
        <div class="table">
          <div class="head">
            <span class="resource">Ресурс</span>
            <span class="date">Дата</span>
          </div>
          <div class="data">
            <div class="row1">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
            <div class="row2">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
            <div class="row3">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
            <div class="row4">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
          </div>
        </div>
      </div>
      <div class="border">
        <div class="line1"></div>
        <div class="line2"></div>
      </div>
    </div>
    <div class="past">
      <div class="title">
        <span class="text">Прошлая неделя</span>
      </div>
      <div class="noRes">
        <b class="head">Неактивные ресурсы</b>
        <div class="table">
          <div class="head">
            <span class="resource">Ресурс</span>
            <span class="date">Дата</span>
          </div>
          <div class="data">
            <div class="row1">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
            <div class="row2">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
            <div class="row3">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
            <div class="row4">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
          </div>
        </div>
      </div>
      <div class="newWAF">
        <b class="head">Новые ресурсы за WAF</b>
        <div class="table">
          <div class="head">
            <span class="resource">Ресурс</span>
            <span class="date">Дата</span>
          </div>
          <div class="data">
            <div class="row1">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
            <div class="row2">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
            <div class="row3">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
            <div class="row4">
              <span class="name">zni.pag.ase.com</span>
              <span class="date">09.10.2023</span>
            </div>
          </div>
        </div>
      </div>
      <div class="border">
        <div class="line1"></div>
        <div class="line2"></div>
      </div>
    </div>
    <button class="downloadReport">
      <img class="downloadImg" src="../../images/download.svg" alt="" />
      <span class="text">Выгрузить полный отчет</span>
    </button>
  </div>
  <div class="Analytics">
    <div class="dynamics">
      <b class="head">Динамика изменения активности ресурсов</b>
      <div class="graf">
        <div class="legend">
          <div class="resources"><span class="text">ресурсы</span></div>
          <div class="waf"><span class="text">за WAF</span></div>
        </div>
        <div class="values"></div>
      </div>
    </div>
    <div class="waf">
      <b class="head">WAF</b>
      <div class="graf">
        <div class="legend">
          <div class="resources"><span class="text">ресурсы</span></div>
          <div class="waf"><span class="text">за WAF</span></div>
        </div>
        <div class="values"></div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import * as httpClient from "../../httpClient.js";
import * as storage from "../../storage.js";
import moment from "moment";

export default defineComponent({
  name: "main",

  data: function () {
    return {
      now: moment().format("DD-MM-YYYY HH:mm"),
      //для главной инфы
      date: null,
      resources: null,
      deactivateResource: null,
      waf: null,
      owners: null,

      //для недельных результатов
      lastNoResCount: null,
      lastNewWafCount: null,
      currNoResCount: null,
      currNewWafCount: null,
      lastNoRes: [],
      lastNewWaf: [],
      currentNoRes: [],
      currentNewWaf: [],

      //для сертификатов
      current: [],
      next: [],
      num: null,
    };
  },

  mounted() {
    // Подставляем имя пользователя
    this.$data.accFIO = storage.get("accFIO");

    // Берем данные для главного меню
    let sendUrl = "/api/general-statistic";
    httpClient.Get(sendUrl).then((response) => {
      storage.set("cache", response.data.cache);
      let resp = JSON.parse(response.data.body);
      this.date = resp.date;
      this.resources = resp.resources;
      this.deactivateResource = resp.deactivateResource;
      this.owners = resp.owners;
      this.waf = resp.waf;
    });

    // берем данные для результатов недели
    sendUrl = "/api/week-statistic";
    httpClient.Get(sendUrl).then((response) => {
      let resp = JSON.parse(response.data.body);
      this.lastNoResCount = resp.last.no_resolve;
      this.lastNewWafCount = resp.last.new_waf;
      this.currNoResCount = resp.current.no_resolve;
      this.currNewWafCount = resp.current.new_waf;
      this.lastNoRes = resp.last.no_res_resource;
      this.lastNewWaf = resp.last.new_waf_resource;
      this.currentNoRes = resp.current.no_res_resource;
      this.currentNewWaf = resp.current.new_waf_resource;
    });

    //берем данные для сертификатов
    sendUrl = "/api/certificates";
    httpClient.Get(sendUrl).then((response) => {
      let resp = JSON.parse(response.data.body);
      this.$data.current = resp.current;
      this.$data.next = resp.next;
      this.num = resp.current.length + resp.next.length;
    });

    //создаем строки для списка ресурсов
    const rowsContainer = document.getElementById("rowsContainer");
    const dataArray = [
      {
        resource: "zni.pag.ase.com",
        ip: "255.255.255.255",
        resolve: "активен",
        waf: "за WAF",
        ssl: "*.soc.mosreg.ru",
        date: "09.10.2023",
      },
      {
        resource: "zni.pag.ase.com",
        ip: "255.255.255.255",
        resolve: "активен",
        waf: "за WAF",
        ssl: "*.soc.mosreg.ru",
        date: "09.10.2023",
      },
      {
        resource: "zni.pag.ase.com",
        ip: "255.255.255.255",
        resolve: "активен",
        waf: "за WAF",
        ssl: "*.soc.mosreg.ru",
        date: "09.10.2023",
      },
    ];
    dataArray.forEach((item) => {
      const rowDiv = document.createElement("div");
      rowDiv.classList.add("rows");
      Object.keys(item).forEach((key) => {
        const span = document.createElement("span");
        span.classList.add(key);
        span.textContent = item[key];
        rowDiv.appendChild(span);
      });
      const lineSpan = document.createElement("span");
      lineSpan.classList.add("line");
      rowDiv.appendChild(lineSpan);
      rowsContainer.appendChild(rowDiv);
    });
  },
});
</script>

<style lang="scss"></style>
