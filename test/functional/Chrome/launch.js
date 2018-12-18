'use strict';

const puppeteer = require('puppeteer');
const config = require('./config.js');

module.export = launch();

async function launch() {
    const browser = await puppeteer.launch(config.puppeteerConfig);

    const page = await browser.newPage();
     
    await page.setViewport(config.viewport);

    await page.goto(config.URL, {waitUntil: 'networkidle2'});

    return page;
}

async function login(page) {
    const loginButton = await page.$('#loginButton');
    
    await loginButton.click();

    await page.waitFor(1000);

    await page.type('#inputCompanyLogin', 'arqueosur');
    await page.type('#inputDepartmentLogin', 'operaciones');
    await page.type('#inputUsernameLogin', 'pedro');
    await page.type('#inputPasswordLogin', 'pass');

    const loginActionButton = await page.$('#loginAction');

    loginActionButton.click();

    await page.waitFor(4000);
}

async function checkUser(page) {
    const userButton = await page.$('#userButton');

    const username = await page.$eval('#userButton', (sel) => {
        return sel.innerHTML;
    });

    if(username == 'pedro'){
        console.log("test checkUser superado");
    } else {
        console.log("test checkUser fallido");
    }
}

async function checkPerlScript(page) {
    const runPerlScript = (await page.$$(".btn.btn-outline-primary"))[2]; //that's not the proper way to select the perl script button

    runPerlScript.click();

    await page.waitFor(2000);

    const runButton = await page.$("button#runButton");

    await runButton.click();

    await page.waitFor(5000);

    const resultBoxText = await page.$eval('#resultBox > p', (selector) => {
        return selector.innerHTML;
    });

    if(resultBoxText == '&gt; 144') {
        console.log("test perl script superado");
    }
}