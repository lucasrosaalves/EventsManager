using EventsManagerPersistence.Domain.Entities;
using Microsoft.AspNetCore.Mvc;
using System;
using System.Collections.Generic;
using System.Linq;

namespace EventsManagerPersistence.Presentation.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class CountriesController : ControllerBase
    {
        private static readonly Country[] Countries = new[]
        {
            new Country("Br", "Brasil", new List<Region>
            {
                new Region("SE", "Sudeste"),
                new Region("S", "Sul"),
                new Region("N", "Norte"),
                 new Region("NE", "Nordeste"),
                 new Region("CO", "Centro-Oeste"),
            }),
        };


        [HttpGet]
        public IActionResult Get()
        {
            return Ok(Countries);
        }

        [HttpGet]
        public IActionResult Get([FromQuery] string name)
        {
            return Ok(Countries.Where(c=> c.Name == name).ToArray());
        }
    }


}
