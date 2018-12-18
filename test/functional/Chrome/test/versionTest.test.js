const puppeteer = require('puppeteer');

const browserDir  = "ws://10.10.0.4:3000"; //the browser server
// const pageDir     = 'http://10.0.0.8'; 
const pageDir     = "http://www.scriptpanel.com" //the page app to search to

const timeouts = {
  general: 30000,
  test: 18000
}

const config = {
  browserWSEndpoint: browserDir,
  headless: true,
  args: ['--no-sandbox', '--disable-setuid-sandbox']
};

let browser;
let page;

beforeAll(async() => {
  jest.setTimeout(timeouts.general);

  browser = await puppeteer.connect(config);
  page = await browser.newPage();

  await page.goto(pageDir);
});

describe('ScriptPanel Main Page', () => {
  afterAll(() => {
    page.close();
  
    browser.close();
  });
  
  test('page title', async () => {
    const title = await page.title();

    await expect(title).toBe('ScriptPanel');
  }, timeouts.test);
});