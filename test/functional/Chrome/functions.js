const puppeteer = require('puppeteer');
const config = require('./config.js');

async function launch(browser) {
    const page = await browser.newPage();
     
    await page.setViewport(config.viewport);

    await page.goto(config.URL, {waitUntil: 'networkidle2'});

    return page;
}

const functions = {
    signup: async (username, department, company, password, page) => {
        const signupButton = await page.$('#signupButton');
    
        await signupButton.click();

        await page.waitFor(1000);

        await page.type('#inputCompany', company);
        await page.type('#inputDepartment', department);
        await page.type('#inputUser', username);
        await page.type('#inputPassword', password);

        const signupActionButton = await page.$('#signupAction');

        await signupActionButton.click();

        await page.waitFor(1000);
    },

    login: async (username, department, company, password, page) => {
        const loginButton = await page.$('#loginButton');
    
        await loginButton.click();

        await page.waitFor(1000);

        await page.type('#inputCompanyLogin', company);
        await page.type('#inputDepartmentLogin', department);
        await page.type('#inputUsernameLogin', username);
        await page.type('#inputPasswordLogin', password);

        const loginActionButton = await page.$('#loginAction');

        await loginActionButton.click();

        await page.waitFor(4000);
    },

    addScript: (script) => {

    },

    runScript: (script) => {

    } 
}

module.exports = functions;


test('Signup and login test', async () => {
    expect.assertions(1);

    const browser = await puppeteer.launch(config.puppeteerConfig);

    const page = await launch(browser);

    await functions.signup("jack", "software", "google", "jack", page);

    await functions.login("jack", "software", "google", "jack", page);

    expect(await page.$eval('#userButton', (selector) => { selector.innerHTML })).toBe("jack");

    browser.close();
}, 16000)