import { Estudiante } from './estudiante';

describe('Estudiante', () => {
  it('should create an instance', () => {
    expect(new Estudiante(1,"","","","",[])).toBeTruthy();
  });
});
