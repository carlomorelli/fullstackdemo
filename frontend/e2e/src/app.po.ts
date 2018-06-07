import { protractor, browser, by, element } from 'protractor';

export class AppPage {
  navigateTo() {
    return browser.get('/');
  }

  getParagraphText() {
    return element(by.css('app-root h1')).getText();
  }

  waitForHomePage() {
    var EC = protractor.ExpectedConditions;
    browser.wait(EC.urlContains('/home'), 5000);
  }

  getLoginBox() {
    return element(by.name('email'))
  }

  getPasswordBox() {
    return element(by.name('password'))
  }

  getSubmitButton() {
    return element(by.buttonText('Login'))
  }

  loginWith(username: string, password: string) {
    this.getLoginBox().sendKeys(username)
    this.getPasswordBox().sendKeys(password)
    this.getSubmitButton().click()
  }

  getHomePageText() {
    return element(by.cssContainingText('p', 'Ciao')).getText()
  }

}
