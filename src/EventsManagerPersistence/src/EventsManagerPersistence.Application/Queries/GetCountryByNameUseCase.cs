using EventsManagerPersistence.Domain.Entities;
using System.Linq;

namespace EventsManagerPersistence.Application.Queries
{
    public class GetCountryByNameUseCase
    {
        private static readonly Country[] _countries = new[]
        {
            new Country("Br", "Brasil"),
        };

        private static readonly Region[] regions = new[]
        {
                new Region("SE", "Sudeste"),
                new Region("S", "Sul"),
                new Region("N", "Norte"),
                new Region("NE", "Nordeste"),
                new Region("CO", "Centro-Oeste"),
        };

        public Country Handle(string countryName)
        {
            return _countries.FirstOrDefault(c => c.Name == countryName);
        }
    }
}
