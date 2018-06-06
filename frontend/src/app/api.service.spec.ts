import { TestBed, inject } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { HttpEvent, HttpEventType } from '@angular/common/http';
import { ApiService } from './api.service';
import { AuthService } from './auth.service';

const token = "someToken"
const backendUrl = "http://localhost:9000"

describe('ApiService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ApiService],
      imports: [ HttpClientTestingModule ]
    });
  });

  it('should be created', inject([ApiService], (service: ApiService) => {
    expect(service).toBeTruthy();
  }));

  it('should retrieve token from AuthService', inject([ApiService], (service: ApiService) => {
    const auth = TestBed.get(AuthService)
    let spy = spyOn(auth, 'getToken').and.returnValue(token)
    expect(service.getHeader().headers.get("Authorization")).toBe(token)
    expect(spy).toHaveBeenCalled();
  }));

  it('should get user detail from HttpClient', inject([ApiService], (service: ApiService) => {
    const auth = TestBed.get(AuthService)
    const backendMock = TestBed.get(HttpTestingController)
    const userId = 123
    const mockUser = { name: 'Bob', surname: 'Ross', age: 41 }

    // patch AuthService and subscribe the stubbed HttpTestingController
    let spy = spyOn(auth, 'getToken').and.returnValue(token)
    service.getUser(userId).subscribe((event: HttpEvent<any>) => {
      switch (event.type) {
        case HttpEventType.Response:
          expect(event.body).toEqual(mockUser);
      }
    });
    var req = backendMock.expectOne(backendUrl + "/users/" + userId)
    expect(req.request.method).toBe("GET")
    expect(req.request.headers.get("Authorization")).toBe(token)
    expect(req.request.responseType).toEqual('json');
    backendMock.verify()
  }));

  it('should not get user detail from HttpClient without token', inject([ApiService], (service: ApiService) => {
    const backendMock = TestBed.get(HttpTestingController)
    const userId = 123
    const mockUser = { name: 'Bob', surname: 'Ross', age: 41 }

    // do not patch AuthService and subscribe the stubbed HttpTestingController
    service.getUser(userId).subscribe((event: HttpEvent<any>) => {
      switch (event.type) {
        case HttpEventType.Response:
          expect(event.body).toEqual(mockUser);
      }
    });
    var req = backendMock.expectOne(backendUrl + "/users/" + userId)
    expect(req.request.method).toBe("GET")
    expect(req.request.body).toBeFalsy()
    backendMock.verify()
  }));


  it('should send login credentials', inject([ApiService], (service: ApiService) => {
    const backendMock = TestBed.get(HttpTestingController)
    let loginResponse
    service.login("someUsername", "somePassword").subscribe((response) => {
      loginResponse = response
    })
    backendMock.expectOne({url: backendUrl + "/login", method: "POST"}).flush("response")
    expect(loginResponse).toEqual("response")
  }));

});
