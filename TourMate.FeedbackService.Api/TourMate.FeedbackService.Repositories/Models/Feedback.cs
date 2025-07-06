using System;
using System.Collections.Generic;

namespace TourMate.FeedbackService.Repositories.Models;

public partial class Feedback
{
    public int FeedbackId { get; set; }

    public int CustomerId { get; set; }

    public int TourGuideId { get; set; }

    public DateTime CreatedDate { get; set; }

    public string Content { get; set; } = null!;

    public int Rating { get; set; }

    public bool IsDeleted { get; set; }

    public DateTime UpdatedAt { get; set; }

    public int InvoiceId { get; set; }
}
