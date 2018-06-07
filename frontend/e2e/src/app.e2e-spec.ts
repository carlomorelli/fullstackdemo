import { AppPage } from './app.po';

describe('workspace-project App', () => {
  let page: AppPage;

  beforeEach(() => {
    page = new AppPage();
    page.navigateTo();
  });


  it('should display welcome message', () => {
    expect(page.getParagraphText()).toEqual('Welcome to empamini!');
  });

  //TC201-1
  it('should have a login and password form', () => {
    expect(page.getLoginBox()).toBeTruthy()
    expect(page.getPasswordBox()).toBeTruthy()
    expect(page.getSubmitButton()).toBeTruthy()
  });

  //TC201-2
  it('should have an error banner if login is not valid', () => {
    page.loginWith("invalidUser", "invalidPass")
  });

  //TC201-2
  it('should have an error banner if password is not valid', () => {
    page.loginWith("demo@emaptica.com", "invalidPass")
  });

  //TC201-3 and TC-202-2
  it('should succesfully login if credentials are ok', () => {
    page.loginWith("demo@empatica.com", "passw0rd")
    page.waitForHomePage() 
    var message = page.getHomePageText()
    expect(message).toContain("Ciao John")
  });

});
