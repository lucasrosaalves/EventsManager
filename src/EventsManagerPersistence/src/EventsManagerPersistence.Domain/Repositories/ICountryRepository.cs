using EventsManagerPersistence.Domain.Entities;
using System.Collections.Generic;

namespace EventsManagerPersistence.Domain.Repositories
{
    public interface ICountryRepository
    {
        IEnumerable<Country> GetAll();

        IEnumerable<Country> FindByName(string name);
    }
}
