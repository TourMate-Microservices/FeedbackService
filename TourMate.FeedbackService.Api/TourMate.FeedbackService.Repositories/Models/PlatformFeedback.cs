using System;
using System.Collections.Generic;

namespace TourMate.FeedbackService.Repositories.Models;

public partial class PlatformFeedback
{
    public int FeedbackId { get; set; }

    public int AccountId { get; set; }

    public int PaymentId { get; set; }

    public int Rating { get; set; }

    public string? Content { get; set; }

    public DateTime CreatedAt { get; set; }
}
