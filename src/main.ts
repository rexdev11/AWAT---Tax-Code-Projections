import puppeteer = require("puppeteer");
import { TEMP_CREDS } from "../.local.env";
import { Page } from "puppeteer";

async function getDocHTML(page: Page): Promise<string> {
    const result = await page
        .evaluate(
            'new XMLSerializer().serializeToString(document.doctype) + document.documentElement.outerHTML;'
        );
    if (
        result
        && typeof result
        === typeof 'string'
    ) {
        return result;
    }
    return await Promise.reject('No Document');
}

export async function sleep(time: number) {
    return await setTimeout(() => {}, time);
}

export async function generalCrawler(options: {
    url: string,
    params: string,
    regEx: RegExp[],
    patterns: string[],
    cb: (err, data) => void
}) {

}

export async function runPoliticalAdCrawler(searchTerm: string): Promise<void> {
    const politicalAdURL: string = `https://www.facebook.com/politicalcontentads/?active_status=all&ad_type=ads-with-political-content&q=${searchTerm}`;

    const Browser = await puppeteer.launch({
        headless: false,
        devtools: true
    });

    const Page = await Browser.newPage();
    await Page.goto(politicalAdURL);
    await Page.type('#' + Object.keys(TEMP_CREDS)[0], TEMP_CREDS.email);
    await Page.type('#' + Object.keys(TEMP_CREDS)[1], TEMP_CREDS.pass);
    await Page.click('#loginbutton');
    await Page.waitForNavigation({
        waitUntil: "networkidle0"
    });

    interface Tagged {
        tags: string[],
        count: number
    }
    const taggedElements: Tagged = await Page.evaluate(() => {
       const anchors: NodeListOf<HTMLAnchorElement> =  document.getElementsByTagName('a');
       const ID_TAG = 'AnchorTagged:';
       const RESULT: {
           tags: string[],
           count: number
       } = {
           tags: [],
           count: 0
       };

       for (let i = 0; i < anchors.length; i++) {
           if (anchors.item(i).className.match('_235y')) {
            anchors.item(i).id = ID_TAG + i;
            RESULT.tags = RESULT.tags.concat(ID_TAG + i)
           }
       }
       RESULT.count = RESULT.tags.length;
       return RESULT;
    });

    for (let tag of taggedElements.tags) {
        await Page.click('#' + tag);
        console.log(tag);
    }

    await sleep(5000);
    console.log('after', await getDocHTML(Page));
    return await Promise.resolve();
}

runPoliticalAdCrawler('Jack');