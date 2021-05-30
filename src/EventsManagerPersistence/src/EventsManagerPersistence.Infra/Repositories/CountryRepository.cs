using EventsManagerPersistence.Domain.Entities;
using EventsManagerPersistence.Domain.Repositories;
using System.Collections.Generic;
using System.Linq;

namespace EventsManagerPersistence.Infra.Repositories
{
    public class CountryRepository : ICountryRepository
    {
        private readonly Country[] _countries;

        public CountryRepository()
        {
            _countries = new[]
            {
                new Country("Br", "Brasil")
            };
        }

        public IEnumerable<Country> GetAll()
        {
            return _countries;
        }

        public IEnumerable<Country> FindByName(string name)
        {
            return _countries.Where(c => c.Name == name);
        }
    }
}
