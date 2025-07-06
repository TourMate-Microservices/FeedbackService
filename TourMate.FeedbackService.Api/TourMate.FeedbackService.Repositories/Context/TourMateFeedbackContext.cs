using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using System;
using System.Collections.Generic;
using TourMate.FeedbackService.Repositories.Models;

namespace TourMate.FeedbackService.Repositories.Context;

public partial class TourMateFeedbackContext : DbContext
{
    public TourMateFeedbackContext()
    {
    }

    public TourMateFeedbackContext(DbContextOptions<TourMateFeedbackContext> options)
        : base(options)
    {
    }

    public virtual DbSet<Feedback> Feedbacks { get; set; }

    public virtual DbSet<PlatformFeedback> PlatformFeedbacks { get; set; }

    public static string GetConnectionString(string connectionStringName)
    {
        var config = new ConfigurationBuilder()
            .SetBasePath(AppDomain.CurrentDomain.BaseDirectory)
            .AddJsonFile("appsettings.json")
            .Build();
        string connectionString = config.GetConnectionString(connectionStringName);
        return connectionString;
    }
    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
       => optionsBuilder.UseSqlServer(GetConnectionString("DefaultConnection"));


    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<Feedback>(entity =>
        {
            entity
                .HasNoKey()
                .ToTable("Feedback");

            entity.Property(e => e.Content).HasColumnName("content");
            entity.Property(e => e.CreatedDate)
                .HasColumnType("datetime")
                .HasColumnName("createdDate");
            entity.Property(e => e.CustomerId).HasColumnName("customerId");
            entity.Property(e => e.FeedbackId)
                .ValueGeneratedOnAdd()
                .HasColumnName("feedbackId");
            entity.Property(e => e.InvoiceId).HasColumnName("invoiceId");
            entity.Property(e => e.IsDeleted).HasColumnName("isDeleted");
            entity.Property(e => e.Rating).HasColumnName("rating");
            entity.Property(e => e.TourGuideId).HasColumnName("tourGuideId");
            entity.Property(e => e.UpdatedAt)
                .HasColumnType("datetime")
                .HasColumnName("updatedAt");
        });

        modelBuilder.Entity<PlatformFeedback>(entity =>
        {
            entity
                .HasNoKey()
                .ToTable("PlatformFeedback");

            entity.Property(e => e.CreatedAt).HasColumnType("datetime");
            entity.Property(e => e.FeedbackId).ValueGeneratedOnAdd();
        });

        OnModelCreatingPartial(modelBuilder);
    }

    partial void OnModelCreatingPartial(ModelBuilder modelBuilder);
}
