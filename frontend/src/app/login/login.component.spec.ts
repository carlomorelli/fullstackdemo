import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { LoginComponent } from './login.component';
import { FormsModule } from '@angular/forms';
import { MatFormFieldModule, MatSnackBarModule, MatInputModule } from '@angular/material';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { when } from 'q';
describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ LoginComponent ],
      imports: [ MatInputModule, MatSnackBarModule, RouterTestingModule, FormsModule, MatFormFieldModule, HttpClientTestingModule, BrowserAnimationsModule ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  // it('should check true when username and password are filled', () => {
  //   component.email = "something"
  //   component.password = "asdasd"
  //   expect(component.check()).toBe(true);
  // });

  // it('should check false when username is not filled', () => {
  //   component.email = ""
  //   component.password = "asdasd"
  //   expect(component.check()).toBe(false);
  // });
  
  // it('should login when credentials are provided', () => {
  //   component.email = "valid"
  //   component.password = "valid"
  //   when(component.api.login())
  // });
});
