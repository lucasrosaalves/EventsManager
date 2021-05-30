using System;

namespace EventsManagerPersistence.Domain.Entities
{
    public class Region
    {
        public Region(string name, string description)
        {
            Id = Guid.NewGuid();
            Name = name;
            Description = description;
        }

        public Guid Id { get; set; }
        public string Name { get; set; }
        public string Description { get; set; }
        public Guid CountryId { get; set; }
    }
}
