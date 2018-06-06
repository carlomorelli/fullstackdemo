import { TestBed, inject } from '@angular/core/testing';

import { AuthService } from './auth.service';

describe('AuthService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [AuthService]
    });
  });

  it('should be created', inject([AuthService], (service: AuthService) => {
    expect(service).toBeTruthy();
  }));

  it('can store a token reference', inject([AuthService], (service: AuthService) => {
    const token = "someToken"
    service.setToken(token)
    expect(service.getToken()).toBe(token);
  }));

  it('can delete a token reference', inject([AuthService], (service: AuthService) => {
    const token = "someOtherToken"
    service.setToken(token)
    service.deleteToken()
    expect(service.getToken()).toBeFalsy();
  }));

});
